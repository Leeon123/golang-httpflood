package stats

import (
	"fmt"
	"time"
)

// DisplayStats prints the current statistics to console
func DisplayStats(snapshot MetricsSnapshot) {
	fmt.Printf("\r\n")
	fmt.Printf("========== STATISTICS ==========\r\n")
	fmt.Printf("Total Requests:      %d\r\n", snapshot.TotalRequests)
	fmt.Printf("Successful:          %d (%.2f%%)\r\n", snapshot.SuccessfulReqs, snapshot.GetSuccessRate())
	fmt.Printf("Failed:              %d (%.2f%%)\r\n", snapshot.FailedReqs, snapshot.GetFailureRate())
	fmt.Printf("Requests/sec:        %.2f\r\n", snapshot.RequestsPerSec)
	fmt.Printf("Active Connections:  %d\r\n", snapshot.ConnectionsOpen)
	fmt.Printf("Bytes Sent:          %d (%.2f MB)\r\n", snapshot.BytesSent, float64(snapshot.BytesSent)/(1024*1024))
	fmt.Printf("Bandwidth:           %.2f MB/s\r\n", snapshot.GetBandwidthMB())
	fmt.Printf("Avg Response Time:   %.2f ms\r\n", snapshot.AvgResponseTimeMs)
	fmt.Printf("Elapsed Time:        %.2f seconds\r\n", snapshot.ElapsedTime)
	fmt.Printf("===============================\r\n")
}

// StartDisplay starts a goroutine that updates statistics display periodically
func (m *Metrics) StartDisplay(interval time.Duration) chan struct{} {
	stopChan := make(chan struct{})
	
	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()
		
		for {
			select {
			case <-ticker.C:
				m.UpdateRequestsPerSecond()
				snapshot := m.GetSnapshot()
				DisplayStats(snapshot)
			case <-stopChan:
				return
			}
		}
	}()
	
	return stopChan
}

// DisplaySummary displays a final summary when the operation completes
func DisplaySummary(snapshot MetricsSnapshot) {
	fmt.Printf("\r\n\r\n")
	fmt.Printf("========== FINAL SUMMARY ==========\r\n")
	fmt.Printf("Total Requests:      %d\r\n", snapshot.TotalRequests)
	fmt.Printf("Successful:          %d\r\n", snapshot.SuccessfulReqs)
	fmt.Printf("Failed:              %d\r\n", snapshot.FailedReqs)
	fmt.Printf("Success Rate:        %.2f%%\r\n", snapshot.GetSuccessRate())
	fmt.Printf("Total Bytes Sent:    %d (%.2f MB)\r\n", snapshot.BytesSent, float64(snapshot.BytesSent)/(1024*1024))
	fmt.Printf("Avg Bandwidth:       %.2f MB/s\r\n", snapshot.GetBandwidthMB())
	fmt.Printf("Total Elapsed Time:  %.2f seconds\r\n", snapshot.ElapsedTime)
	fmt.Printf("==================================\r\n")
}
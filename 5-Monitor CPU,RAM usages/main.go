package main

import (
	"log"
	"net/http"
	"runtime"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	cpuUsage = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "app_cpu_usage_percent",
		Help: "CPU usage of the app",
	})
	memUsage = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "app_memory_usage_bytes",
		Help: "Memory usage of the app in bytes",
	})
)

func init() {
	// Register the metrics
	prometheus.MustRegister(cpuUsage)
	prometheus.MustRegister(memUsage)
}

// Dummy function to simulate CPU usage measurement
func measureCPUUsage() float64 {
	// CPU measurement per Go process is tricky; here we simulate or stub it
	// You can integrate actual CPU usage with gopsutil or cgroups for container-level stats
	return float64(runtime.NumGoroutine()) * 1.5 // Fake value for example
}

func recordMetrics() {
	go func() {
		for {
			var m runtime.MemStats
			runtime.ReadMemStats(&m)

			memUsage.Set(float64(m.Alloc))  // RAM used by Go app
			cpuUsage.Set(measureCPUUsage()) // Simulated CPU usage

			time.Sleep(5 * time.Second) // Update interval
		}
	}()
}

func main() {
	recordMetrics()

	http.Handle("/metrics", promhttp.Handler())
	log.Println("✅ Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

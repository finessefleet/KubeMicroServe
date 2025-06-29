// main.go: Entry point for the KubeMicroServe Go web server
// Sets up HTTP routes, integrates Prometheus metrics, and starts the server
package main

import (
	"fmt"
	"net/http"

	"basic-server/internal/handlers"
	"basic-server/internal/middleware"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// main initializes HTTP handlers and starts the web server
func main() {
	// Register the root endpoint with metrics middleware
	http.HandleFunc("/", middleware.Metrics(handlers.Home))
	// Register the health check endpoint with metrics middleware
	http.HandleFunc("/health", middleware.Metrics(handlers.Health))
	// Expose Prometheus metrics endpoint
	http.Handle("/metrics", promhttp.Handler())

	fmt.Println("Server starting on :8080...")
	// Start the HTTP server on port 8080
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(fmt.Sprintf("Server failed to start: %v", err))
	}
}
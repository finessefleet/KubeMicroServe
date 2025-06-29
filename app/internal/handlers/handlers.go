package handlers

import (
	"fmt"
	"net/http"
)

// handlers.go: HTTP handler functions for KubeMicroServe
// Defines endpoints for home and health checks

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Request received from IP: %s\n", r.RemoteAddr)
	fmt.Fprintln(w, "Hello from Go!")
}

func Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "OK")
}
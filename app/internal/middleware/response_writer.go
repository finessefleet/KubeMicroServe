package middleware

import "net/http"

// response_writer.go: Custom HTTP response writer for KubeMicroServe
// Used to capture status codes for Prometheus metrics

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func NewResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{w, http.StatusOK} 
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func (rw *responseWriter) StatusCode() int {
	return rw.statusCode
}
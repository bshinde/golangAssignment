package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestLatencyLogger(t *testing.T) {
	// Create a test handler to check middleware
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	// Wrap the handler with the LatencyLogger middleware
	mw := LatencyLogger(handler)

	// Create a new HTTP request
	req := httptest.NewRequest("GET", "/test", nil)
	w := httptest.NewRecorder()

	// Capture start time
	startTime := time.Now()

	// Serve the HTTP request
	mw.ServeHTTP(w, req)

	// Ensure the request completes and logs the latency
	duration := time.Since(startTime)

	// Assert that the duration is greater than 0 (meaning it took some time)
	if duration <= 0 {
		t.Errorf("Expected request duration to be greater than 0, but got %v", duration)
	}

	// Assert the status code and body
	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200 OK, got %d", w.Code)
	}
	if w.Body.String() != "OK" {
		t.Errorf("Expected response body 'OK', got %s", w.Body.String())
	}
}

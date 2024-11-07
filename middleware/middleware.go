package middleware

import (
	"log"
	"net/http"
	"time"
)

// LatencyLogger is a middleware that logs the time taken to process each request
func LatencyLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		duration := time.Since(start)
		log.Printf("Request %s %s took %v", r.Method, r.URL.Path, duration)
	})
}

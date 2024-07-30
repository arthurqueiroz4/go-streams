package middleware

import (
	"fmt"
	"net/http"
	"time"
)

var TimingMid = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		duration := time.Since(start)
		fmt.Printf("Request processed in %v\n", duration)
	})
}

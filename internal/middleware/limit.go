package middleware

import (
	"net/http"

	"golang.org/x/time/rate"
)

var limiter = rate.NewLimiter(20, 10)

func Limit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !limiter.Allow() {
			http.Error(w, http.StatusText(429), http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, r)
	})
}

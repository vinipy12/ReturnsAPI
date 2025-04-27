package main

import (
	"log"
	"net/http"
	"time"

	"github.com/vinipy12/ReturnsAPI/ratelimiter"
)

type Middleware func(http.Handler) http.HandlerFunc

func MiddlewareChain(middlewares ...Middleware) Middleware {
	return func(next http.Handler) http.HandlerFunc {
		for i := len(middlewares) - 1; i >= 0; i-- {
			next = middlewares[i](next)
		}

		return next.ServeHTTP
	}
}

func requestLoggerMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		formattedTime := time.Now().Format("2006-01-02 15:04:05.00")
		log.Printf("%s - %s - %s", r.Method, r.URL.Path, formattedTime)
		next.ServeHTTP(w, r)
	}
}

// For extracting IP, can use r.RemoteAddr, but for production could use a more robust solution like X-Forwarded-For header
// Testing with curl makes IP Address looks weird, like [::1]:54321, and each request has a different port, thus, if not parsed correctly, each request will be a new IP:PORT
func parseIpAddress(ip string) string {
	parsedIp := ip[:5]
	return parsedIp
}

func rateLimiterMiddleware(next http.Handler) http.HandlerFunc {
	rateLimiter := ratelimiter.InMemoryRateLimiter
	return func(w http.ResponseWriter, r *http.Request) {
		if !rateLimiter.AllowRequest(parseIpAddress(r.RemoteAddr)) {
			w.WriteHeader(http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	}
}

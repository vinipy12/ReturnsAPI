package main

import (
	"log"
	"net/http"
	"time"
)

func requestLoggerMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s - %s - %s", r.Method, r.URL.Path, time.Now())
		next.ServeHTTP(w, r)
	}
}

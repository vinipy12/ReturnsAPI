package main

import (
	"log"
	"net/http"
	"time"
)

func requestLoggerMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		formattedTime := time.Now().Format("2006-01-02 15:04:05.00")
		log.Printf("%s - %s - %s", r.Method, r.URL.Path, formattedTime)
		next.ServeHTTP(w, r)
	}
}

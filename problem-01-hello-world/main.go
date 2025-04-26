package main

import "net/http"

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("message: Hello, World!"))
	})
	http.ListenAndServe(":8080", nil)
}

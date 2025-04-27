package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
)

type ApiServer struct {
	addr string
}

func newApiServer(addr string) *ApiServer {
	return &ApiServer{
		addr: addr,
	}
}

type returnRequest struct {
	OrderId string `json:"orderId"`
}

var allRequests = make(map[uuid.UUID]returnRequest)

func validatePayload(payload returnRequest) bool {
	if payload.OrderId == "" || len(payload.OrderId) > 100 {
		return false
	}
	return true
}

func newReturn(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	if r.Header.Get("Content-Type") != "application/json" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		return
	}

	var payload returnRequest
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	if err := decoder.Decode(&payload); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !validatePayload(payload) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id := uuid.New()
	allRequests[id] = payload

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		log.Printf("Failed to encode response: %v", err)
	}
}

func getReturn(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	strId := r.PathValue("id")
	parsedId, err := uuid.Parse(strId)
	if err != nil {
		log.Printf("Failed to parse id: %v", err)
	}

	encoder := json.NewEncoder(w)
	order := allRequests[parsedId]
	w.WriteHeader(http.StatusOK)
	if err := encoder.Encode(order); err != nil {
		log.Printf("Failed to encode response: %v", err)
	}

}

func (s *ApiServer) Run() error {
	router := http.NewServeMux()

	router.HandleFunc("POST /returns", newReturn)
	router.HandleFunc("GET /returns/{id}", getReturn)

	middlewareChain := MiddlewareChain(
		requestLoggerMiddleware,
		rateLimiterMiddleware,
	)

	server := http.Server{
		Addr:    s.addr,
		Handler: middlewareChain(router),
	}

	return server.ListenAndServe()
}

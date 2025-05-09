package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/vinipy12/ReturnsAPI/logger"
)

type ApiServer struct {
	addr   string
	logger logger.Logger
}

func newApiServer(addr string) *ApiServer {
	return &ApiServer{
		addr:   addr,
		logger: *logger.NewLogger(),
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
		sendError(w, http.StatusMethodNotAllowed, "Invalid HTTP Method")
		return
	}

	if r.Header.Get("Content-Type") != "application/json" {
		sendError(w, http.StatusUnsupportedMediaType, "Invalid Content-Type Header")
		return
	}

	var payload returnRequest
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	if err := decoder.Decode(&payload); err != nil {
		sendError(w, http.StatusBadRequest, "Failed to decode payload")
		return
	}

	if !validatePayload(payload) {
		sendError(w, http.StatusBadRequest, "Invalid Payload")
		return
	}

	id := uuid.New()
	allRequests[id] = payload

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		logger.Log(fmt.Sprintf("Failed to encode response: %v", err))
	}
}

func getReturn(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		sendError(w, http.StatusMethodNotAllowed, "Invalid HTTP Method")
		return
	}

	strId := r.PathValue("id")
	parsedId, err := uuid.Parse(strId)
	if err != nil {
		logger.Log(fmt.Sprintf("Failed to parse id: %v", err))
	}

	encoder := json.NewEncoder(w)
	order := allRequests[parsedId]
	w.WriteHeader(http.StatusOK)
	if err := encoder.Encode(order); err != nil {
		logger.Log(fmt.Sprintf("Failed to encode response: %v", err))
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

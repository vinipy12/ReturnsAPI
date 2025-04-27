package main

import (
	"encoding/json"
	"fmt"
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

func returnsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var req returnRequest
		decoder := json.NewDecoder(r.Body)
		defer r.Body.Close()

		if err := decoder.Decode(&req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if req.OrderId == "" || len(req.OrderId) > 100 {
			w.WriteHeader(http.StatusBadRequest)
		}

		id := uuid.New()
		allRequests[id] = req
		fmt.Println(id)

		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(req); err != nil {
			log.Printf("Failed to encode response: %v", err)
		}
	} else {
		strId := r.PathValue("id")
		parsedId, err := uuid.Parse(strId)
		if err != nil {
			log.Printf("Failed to parse id: %v", err)
		}

		order := allRequests[parsedId]
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(order); err != nil {
			log.Printf("Failed to encode response: %v", err)
		}
	}

}

func (s *ApiServer) Run() error {
	router := http.NewServeMux()

	router.HandleFunc("POST /returns", returnsHandler)
	router.HandleFunc("GET /returns/{id}", returnsHandler)

	server := http.Server{
		Addr:    s.addr,
		Handler: router,
	}

	return server.ListenAndServe()
}

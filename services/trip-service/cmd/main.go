package main

import (
	"log"
	"net/http"

	h "github.com/erickgreco/micro-ride/services/trip-service/internal/infrastructure/http"
	"github.com/erickgreco/micro-ride/services/trip-service/internal/infrastructure/repository"
	"github.com/erickgreco/micro-ride/services/trip-service/internal/service"
)

func main() {
	inMemRepo := repository.NewInmemRepository()
	srvc := service.NewService(inMemRepo)
	mux := http.NewServeMux()

	httpHandler := h.HttpHandler{Service: srvc}

	mux.HandleFunc("POST /preview", httpHandler.HandleTripPreview)

	server := &http.Server{
		Addr:    ":8083",
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Printf("HTTP server error: %v", err)
	}
}

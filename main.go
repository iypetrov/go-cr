package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/iypetrov/go-cr/config"
	"github.com/iypetrov/go-cr/logger"
	"github.com/iypetrov/go-cr/registry"
)

func main() {
	mux := chi.NewRouter()

	cfg := config.New()
	log := logger.New(cfg)

	registry := registry.New(log)

	mux.With().Route("/v2", registry.Router())

	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", cfg.App.Port),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
		Handler:      mux,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Error("cannot start server: %s", err.Error())
	}
}

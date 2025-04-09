package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	awscfg "github.com/aws/aws-sdk-go-v2/config"

	"github.com/go-chi/chi/v5"
	"github.com/iypetrov/go-cr/config"
	"github.com/iypetrov/go-cr/distribution"
	"github.com/iypetrov/go-cr/logger"
)

func main() {
	mux := chi.NewRouter()

	ctx := context.Background()
	cfg := config.New()
	log := logger.New(cfg)
	awsCfg, err := awscfg.LoadDefaultConfig(
		ctx,
		awscfg.WithRegion(cfg.AWS.Region),
	)
	if err != nil {
		log.Error("Failed to load AWS config: %v", err.Error())
	}

	storage := distribution.NewStorage(awsCfg, log)
	metadata := distribution.NewMetadata(awsCfg, log)
	registry := distribution.NewRegistry(
		storage,
		metadata,
		log,
	)

	mux.With().Route("/v2", distribution.Router(registry))
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

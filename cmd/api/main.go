package main

import (
	"context"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/alonsoF100/shotlink/internal/config"
	"github.com/alonsoF100/shotlink/internal/logger"
	"github.com/alonsoF100/shotlink/internal/repository/storage/postgres"
	"github.com/alonsoF100/shotlink/internal/service"
	"github.com/alonsoF100/shotlink/internal/transport/http/routing"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Config error: %s", err.Error())
	}

	logger.Setup(cfg.Log)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	pool, err := postgres.NewPool(ctx, cfg.Database)
	if err != nil {
		slog.Error("Failed to connect to database", "error", err)
		return
	}
	defer pool.Close()

	service := service.New(nil, nil)
	router := routing.SetupRouter(service, cfg.Server.BaseURL)

	done := make(chan struct{})

	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
		<-sigChan

		slog.Info("Shutting down gracefully...")
		cancel()
		close(done)
	}()

	go func() {
		slog.Info("Starting server", "port", cfg.Server.Port)
		if err := router.Run(cfg.Server.Addr()); err != nil {
			slog.Error("Server failed", "error", err)
		}
		close(done)
	}()

	<-done
	slog.Info("Shutdown complete")
}

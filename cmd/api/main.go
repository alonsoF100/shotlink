package main

import (
	"log"

	"github.com/alonsoF100/shotlink/internal/config"
	"github.com/alonsoF100/shotlink/internal/logger"
	"github.com/alonsoF100/shotlink/internal/transport/http/handler"
	"github.com/alonsoF100/shotlink/internal/transport/http/routing"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Config error: %s", err.Error())
	}

	logger.Setup(cfg.Log)

	var service handler.Service = nil
	router := routing.SetupRouter(service, cfg.Server.BaseURL)
	router.Run(cfg.Server.Addr())
}

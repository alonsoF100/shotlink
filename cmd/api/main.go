package main

import (
	"fmt"
	"log"

	"github.com/alonsoF100/shotlink/internal/config"
	"github.com/alonsoF100/shotlink/internal/logger"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Config error: %s", err.Error())
	}
	fmt.Println(cfg.Server.Port)

	logger := logger.Setup(cfg.Log)
	logger.Info("dfssf")

}

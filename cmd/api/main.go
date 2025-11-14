package main

import (
	"fmt"
	"log"

	"github.com/alonsoF100/shotlink/internal/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Config error: %s", err.Error())
	}

	fmt.Println(cfg.Server.Port)
}

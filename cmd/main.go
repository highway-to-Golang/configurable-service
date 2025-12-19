package main

import (
	"log/slog"
	"os"

	"highway-to-Golang/configurable-service/internal/config"
	"highway-to-Golang/configurable-service/internal/service"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		slog.Error("Failed to load configuration", "error", err)
		os.Exit(1)
	}

	configService := service.NewConfigService(cfg)
	configService.DisplayConfig()
}

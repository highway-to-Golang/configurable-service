package service

import (
	"log/slog"

	"highway-to-Golang/configurable-service/internal/config"
)

type ConfigService struct {
	config *config.Config
}

func NewConfigService(cfg *config.Config) *ConfigService {
	return &ConfigService{
		config: cfg,
	}
}

func (s *ConfigService) DisplayConfig() {
	slog.Info("=== Configuration Service ===")
	slog.Info("Application Configuration",
		"name", s.config.App.Name,
		"version", s.config.App.Version,
		"environment", s.config.App.Environment,
	)

	slog.Info("Database Configuration",
		"host", s.config.Database.Host,
		"port", s.config.Database.Port,
		"username", s.config.Database.Username,
		"database", s.config.Database.Database,
		"dsn", s.config.GetDatabaseDSN(),
	)

	slog.Info("Logging Configuration",
		"level", s.config.Logging.Level,
		"format", s.config.Logging.Format,
	)
}

func (s *ConfigService) GetConfig() *config.Config {
	return s.config
}

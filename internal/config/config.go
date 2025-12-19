package config

import (
	"errors"
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	App      AppConfig      `mapstructure:"app"`
	Database DatabaseConfig `mapstructure:"database"`
	Logging  LoggingConfig  `mapstructure:"logging"`
}

type AppConfig struct {
	Name        string `mapstructure:"name"`
	Version     string `mapstructure:"version"`
	Environment string `mapstructure:"environment"`
}

type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
}

type LoggingConfig struct {
	Level  string `mapstructure:"level"`
	Format string `mapstructure:"format"`
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		slog.Warn("No .env file found, using environment variables only", "error", err)
	}

	viper.SetDefault("app.name", "configurable-service")
	viper.SetDefault("app.version", "1.0.0")
	viper.SetDefault("app.environment", "development")
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", 5432)
	viper.SetDefault("database.username", "postgres")
	viper.SetDefault("database.password", "")
	viper.SetDefault("database.database", "configurable_service")
	viper.SetDefault("logging.level", "info")
	viper.SetDefault("logging.format", "json")

	viper.SetEnvPrefix("MYAPP")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	viper.AutomaticEnv()

	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./configs")

	if err := viper.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if !errors.As(err, &configFileNotFoundError) {
			return nil, fmt.Errorf("error reading config file: %w", err)
		}
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("error unmarshaling config: %w", err)
	}

	overrideFromEnv(&config)

	return &config, nil
}
func overrideFromEnv(config *Config) {
	if name := os.Getenv("MYAPP_APP_NAME"); name != "" {
		config.App.Name = name
	}
	if version := os.Getenv("MYAPP_APP_VERSION"); version != "" {
		config.App.Version = version
	}
	if env := os.Getenv("MYAPP_APP_ENVIRONMENT"); env != "" {
		config.App.Environment = env
	}
	if host := os.Getenv("MYAPP_DATABASE_HOST"); host != "" {
		config.Database.Host = host
	}
	if port := os.Getenv("MYAPP_DATABASE_PORT"); port != "" {
		if p, err := strconv.Atoi(port); err == nil {
			config.Database.Port = p
		}
	}
	if username := os.Getenv("MYAPP_DATABASE_USERNAME"); username != "" {
		config.Database.Username = username
	}
	if password := os.Getenv("MYAPP_DATABASE_PASSWORD"); password != "" {
		config.Database.Password = password
	}
	if database := os.Getenv("MYAPP_DATABASE_DATABASE"); database != "" {
		config.Database.Database = database
	}
	if level := os.Getenv("MYAPP_LOGGING_LEVEL"); level != "" {
		config.Logging.Level = level
	}
	if format := os.Getenv("MYAPP_LOGGING_FORMAT"); format != "" {
		config.Logging.Format = format
	}
}

func (c *Config) GetDatabaseDSN() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		c.Database.Host, c.Database.Port, c.Database.Username, c.Database.Password, c.Database.Database)
}

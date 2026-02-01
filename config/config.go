package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	DatabaseURL string `mapstructure:"DATABASE_URL"`
	ServerPort  string `mapstructure:"SERVER_PORT"`
	AppEnv      string `mapstructure:"APP_ENV"`
}

func LoadConfig() *Config {
	// Bind specific environment variables (required for Railway/Zeabur)
	viper.BindEnv("DATABASE_URL")
	viper.BindEnv("SERVER_PORT")
	viper.BindEnv("APP_ENV")

	// Try to read .env file (for local development)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	// Set defaults
	viper.SetDefault("SERVER_PORT", "8080")
	viper.SetDefault("APP_ENV", "development")

	if err := viper.ReadInConfig(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalf("Unable to parse config: %v", err)
	}

	// Fallback: check if DATABASE_URL is still empty, read from os.Getenv directly
	if config.DatabaseURL == "" {
		config.DatabaseURL = os.Getenv("DATABASE_URL")
	}
	if config.ServerPort == "" {
		config.ServerPort = os.Getenv("SERVER_PORT")
		// Also try PORT (used by Railway, Zeabur, Heroku, etc.)
		if config.ServerPort == "" {
			config.ServerPort = os.Getenv("PORT")
		}
	}
	if config.AppEnv == "" {
		config.AppEnv = os.Getenv("APP_ENV")
	}

	// Apply defaults if still empty
	if config.ServerPort == "" {
		config.ServerPort = "8080"
	}
	if config.AppEnv == "" {
		config.AppEnv = "development"
	}

	return &config
}

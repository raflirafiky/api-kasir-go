package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DatabaseURL string `mapstructure:"DATABASE_URL"`
	ServerPort  string `mapstructure:"SERVER_PORT"`
	AppEnv      string `mapstructure:"APP_ENV"`
}

func LoadConfig() *Config {
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

	return &config
}

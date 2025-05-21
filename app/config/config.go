package config

import (
	"fmt"
	"log"
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Environment string

type Config struct {
	Environment Environment `required:"true" envconfig:"ENVIRONMENT"`

	App    App
	Server Server

	WeatherAPI WeatherAPI
}

type App struct {
	Name                    string        `required:"true" envconfig:"APP_NAME"`
	ID                      string        `required:"true" envconfig:"APP_ID"`
	GracefulShutdownTimeout time.Duration `required:"true" envconfig:"APP_GRACEFUL_SHUTDOWN_TIMEOUT"`
}

type Server struct {
	Address      string        `required:"true" envconfig:"SERVER_ADDRESS"`
	ReadTimeout  time.Duration `required:"true" envconfig:"SERVER_READ_TIMEOUT"`
	WriteTimeout time.Duration `required:"true" envconfig:"SERVER_WRITE_TIMEOUT"`
}

type WeatherAPI struct {
	URL    string `required:"true" envconfig:"WEATHER_API_URL"`
	APIKey string `required:"true" envconfig:"WEATHER_API_KEY"`
}

func New() (Config, error) {
	const operation = "Config.New"

	var cfg Config

	if err := godotenv.Load(); err != nil {
		log.Println(".env file not found, using system environment variables")
	}

	err := envconfig.Process("", &cfg)
	if err != nil {
		return Config{}, fmt.Errorf("%s -> %w", operation, err)
	}

	return cfg, nil
}

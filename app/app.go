package app

import (
	"context"

	"github.com/leonardo-gmuller/go-weather/app/config"
	"github.com/leonardo-gmuller/go-weather/app/domain/usecase"
)

type App struct {
	UseCase usecase.UseCaseInterface
}

func New(ctx context.Context, config config.Config) (*App, error) {
	usecase := usecase.New(&config)
	return &App{
		UseCase: usecase,
	}, nil
}

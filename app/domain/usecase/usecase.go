package usecase

import (
	"context"

	"github.com/leonardo-gmuller/go-weather/app/config"
	"github.com/leonardo-gmuller/go-weather/app/domain/dto"
	"github.com/leonardo-gmuller/go-weather/app/gateway/client"
)

type UseCaseInterface interface {
	GetAddress(ctx context.Context, zipcode string) (*AddressResponse, error)
	GetWeather(ctx context.Context, address dto.Address) (*WeatherResponse, error)
}

type UseCase struct {
	AppName string

	AddressGateway client.AddressGateway
	WeatherGateway client.WeatherGateway
}

func New(config *config.Config) *UseCase {
	return &UseCase{
		AppName: config.App.Name,

		AddressGateway: client.NewAddressGateway(),
		WeatherGateway: client.NewWeatherGateway(config),
	}
}

package usecase

import (
	"context"

	"github.com/leonardo-gmuller/go-weather/app/domain/dto"
)

type WeatherResponse struct {
	TempC float64
	TempF float64
	TempK float64
}

func (u *UseCase) GetWeather(ctx context.Context, address dto.Address) (*WeatherResponse, error) {
	data, err := u.WeatherGateway.GetWeatherByCity(address.City, address.UF)
	if err != nil {
		return nil, err
	}

	c := data.TempC
	return &WeatherResponse{
		TempC: c,
		TempF: c*1.8 + 32,
		TempK: c + 273,
	}, nil
}

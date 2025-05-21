package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/leonardo-gmuller/go-weather/app/domain/dto"
	"github.com/leonardo-gmuller/go-weather/app/gateway/client"
	"github.com/stretchr/testify/assert"
)

type MockWeatherGateway struct {
	GetWeatherByCityFunc func(city, uf string) (*client.Weather, error)
}

func (m *MockWeatherGateway) GetWeatherByCity(city, uf string) (*client.Weather, error) {
	return m.GetWeatherByCityFunc(city, uf)
}

func TestGetWeather_Success(t *testing.T) {
	mockGateway := &MockWeatherGateway{
		GetWeatherByCityFunc: func(city, uf string) (*client.Weather, error) {
			return &client.Weather{TempC: 25.0}, nil
		},
	}
	u := &UseCase{WeatherGateway: mockGateway}
	address := dto.Address{City: "Sao Paulo", UF: "SP"}

	resp, err := u.GetWeather(context.Background(), address)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, 25.0, resp.TempC)
	assert.InDelta(t, 77.0, resp.TempF, 0.01)
	assert.InDelta(t, 298.0, resp.TempK, 0.01)
}

func TestGetWeather_ErrorFromGateway(t *testing.T) {
	mockGateway := &MockWeatherGateway{
		GetWeatherByCityFunc: func(city, uf string) (*client.Weather, error) {
			return nil, errors.New("gateway error")
		},
	}
	u := &UseCase{WeatherGateway: mockGateway}
	address := dto.Address{City: "Nowhere", UF: "XX"}

	resp, err := u.GetWeather(context.Background(), address)
	assert.Error(t, err)
	assert.Nil(t, resp)
}

func TestGetWeather_ZeroTemperature(t *testing.T) {
	mockGateway := &MockWeatherGateway{
		GetWeatherByCityFunc: func(city, uf string) (*client.Weather, error) {
			return &client.Weather{TempC: 0.0}, nil
		},
	}
	u := &UseCase{WeatherGateway: mockGateway}
	address := dto.Address{City: "ZeroCity", UF: "ZZ"}

	resp, err := u.GetWeather(context.Background(), address)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, 0.0, resp.TempC)
	assert.InDelta(t, 32.0, resp.TempF, 0.01)
	assert.InDelta(t, 273.0, resp.TempK, 0.01)
}

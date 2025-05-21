package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/leonardo-gmuller/go-weather/app/config"
)

type Weather struct {
	TempC float64
}

type WeatherGateway interface {
	GetWeatherByCity(city, uf string) (*Weather, error)
}

func NewWeatherGateway(config *config.Config) *weatherClient {
	return &weatherClient{
		apiKey: config.WeatherAPI.APIKey,
		url:    config.WeatherAPI.URL,
	}
}

type weatherClient struct {
	apiKey string
	url    string
}

type weatherAPIResponse struct {
	Current struct {
		TempC float64 `json:"temp_c"`
	} `json:"current"`
}

func (c *weatherClient) GetWeatherByCity(city, uf string) (*Weather, error) {
	location := fmt.Sprintf("%s,%s", city, uf)
	escapedLocation := url.QueryEscape(location)

	fullURL := fmt.Sprintf("%s?key=%s&q=%s", c.url, c.apiKey, escapedLocation)

	resp, err := http.Get(fullURL)
	if err != nil {
		return nil, fmt.Errorf("error fetching weather: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("weather api error: %s", resp.Status)
	}

	var data weatherAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf("error decoding weather json: %w", err)
	}

	return &Weather{
		TempC: data.Current.TempC,
	}, nil
}

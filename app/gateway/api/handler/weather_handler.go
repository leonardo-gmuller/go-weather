package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/leonardo-gmuller/go-weather/app/domain/usecase"
)

const weatherPattern = "/weather"

func (h *Handler) WeatherSetup(router chi.Router) {
	router.Route(weatherPattern, func(r chi.Router) {
		r.Get("/{cep}", h.GetWeather)
	})
}

func (h *Handler) GetWeather(w http.ResponseWriter, r *http.Request) {
	cep := chi.URLParam(r, "cep")

	address, err := h.useCase.GetAddress(r.Context(), cep)
	if err != nil {
		var status int
		switch err {
		case usecase.ErrInvalidZipcode:
			status = http.StatusUnprocessableEntity
		case usecase.ErrNotFound:
			status = http.StatusNotFound
		default:
			status = http.StatusInternalServerError
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return

	}

	weather, err := h.useCase.GetWeather(r.Context(), address.Address)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := map[string]float64{
		"temp_C": weather.TempC,
		"temp_F": weather.TempF,
		"temp_K": weather.TempK,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

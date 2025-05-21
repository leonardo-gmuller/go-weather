package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const urlViaCep = "https://viacep.com.br/ws/%s/json"

type ViaCepResponse struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
	Erro        bool   `json:"erro,omitempty"`
}

type AddressGateway interface {
	GetAddressByCEP(cep string) (*ViaCepResponse, error)
}

func NewAddressGateway() AddressGateway {
	return &addressGateway{}
}

type addressGateway struct{}

func (a *addressGateway) GetAddressByCEP(cep string) (*ViaCepResponse, error) {
	req, err := http.Get(fmt.Sprintf(urlViaCep, cep))
	if err != nil {
		return &ViaCepResponse{}, err
	}

	defer req.Body.Close()
	body, err := io.ReadAll(req.Body)

	if err != nil {
		return &ViaCepResponse{}, err
	}

	var data ViaCepResponse

	err = json.Unmarshal(body, &data)

	if err != nil {
		return &ViaCepResponse{}, err
	}

	if data.Erro {
		return &ViaCepResponse{}, fmt.Errorf("zipcode not found")
	}

	return &data, nil
}

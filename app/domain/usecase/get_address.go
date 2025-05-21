package usecase

import (
	"context"
	"errors"

	"github.com/leonardo-gmuller/go-weather/app/domain/dto"
)

var (
	ErrInvalidZipcode = errors.New("invalid zipcode")
	ErrNotFound       = errors.New("can not find zipcode")
)

type AddressResponse struct {
	Address dto.Address
}

func (u *UseCase) GetAddress(ctx context.Context, zipcode string) (*AddressResponse, error) {
	if len(zipcode) != 8 {
		return nil, ErrInvalidZipcode
	}

	address, err := u.AddressGateway.GetAddressByCEP(zipcode)
	if err != nil {
		return nil, ErrNotFound
	}

	return &AddressResponse{Address: dto.Address{
		City: address.Localidade,
		UF:   address.Uf,
	}}, nil
}

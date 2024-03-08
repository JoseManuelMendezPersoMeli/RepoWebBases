package service

import (
	"GoWeb/internal"
	"errors"
)

var (
	ErrNameRequired       = errors.New("name is required")
	ErrQuantityRequired   = errors.New("quantity is required or can't be less than 0")
	ErrCodeValueRequired  = errors.New("CodeValue is required")
	ErrExpirationRequired = errors.New("expiration is required")
	ErrPriceRequired      = errors.New("price is required or can't be less than 0")
)

func ValidateKeys(body internal.ProductRequestBody) error {
	if body.Name == "" {
		return ErrNameRequired
	}

	if body.Quantity <= 0 {
		return ErrQuantityRequired
	}

	if body.CodeValue == "" {
		return ErrCodeValueRequired
	}

	if body.Expiration == "" {
		return ErrExpirationRequired
	}

	if body.Price <= 0 {
		return ErrPriceRequired
	}
	return nil
}

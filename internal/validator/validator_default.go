package validator

import (
	"time"
)

func ValidateProductDetails(product *ValidatorProductDetails) error {
	if product.Name == "" {
		return ErrNoName
	}
	if product.Quantity < 0 {
		return ErrNoQuantity
	}
	if product.CodeValue == "" {
		return ErrNoCodeValue
	}
	if time.Time.IsZero(product.Expiration) {
		return ErrNoExpiration
	}
	if product.Price < 0 {
		return ErrNoPrice
	}
	return nil
}

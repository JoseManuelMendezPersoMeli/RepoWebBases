package validator

import (
	"errors"
	"time"
)

type ValidatorProductDetails struct {
	Name        string
	Quantity    int
	CodeValue   string
	IsPublished bool
	Expiration  time.Time
	Price       float64
}

type Validator interface {
	ValidateProductDetails(product *ValidatorProductDetails) error
}

var (
	// ErrNoName This error is returned when the name is not provided
	ErrNoName = errors.New("name is required")

	// ErrNoQuantity This error is returned when the quantity is not provided
	ErrNoQuantity = errors.New("quantity is required")

	// ErrNoCodeValue This error is returned when the code value is not provided
	ErrNoCodeValue = errors.New("code value is required")

	// ErrNoExpiration This error is returned when the expiration is not provided
	ErrNoExpiration = errors.New("expiration is required")

	// ErrNoPrice This error is returned when the price is not provided
	ErrNoPrice = errors.New("price is required")
)

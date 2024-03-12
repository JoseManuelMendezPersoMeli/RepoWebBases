package storage

import (
	"errors"
	"time"
)

type Product struct {
	ID          int
	Name        string
	Quantity    int
	CodeValue   string
	IsPublished bool
	Expiration  time.Time
	Price       float64
}

type Query struct {
	PriceGt float64
}

type StoredProduct interface {
	// Ping returns a pong message
	Ping() (string, error)

	// GetAll return all products from the warehouse
	GetAll() ([]*Product, error)

	// GetByID return a product by ID from the warehouse
	GetByID(id int) (*Product, error)

	// Search returns a list of products that match the search criteria
	Search(query *Query) ([]*Product, error)

	// AddProduct adds a new product to the warehouse
	AddProduct(product *ProductAttributesDefault) error

	// UpdateOrCreateProduct updates or creates a product in the warehouse
	UpdateOrCreateProduct(id int, product *ProductAttributesDefault) error

	// UpdateProduct updates a product in the warehouse
	UpdateProduct(id int, product map[string]any) (*Product, error)

	// DeleteProduct deletes a product from the warehouse
	DeleteProduct(id int) (*Product, error)
}

var (
	// ErrDatabaseAccess This error is returned when there is an error accessing the database
	ErrDatabaseAccess = errors.New("error accessing the database")

	// ErrProductNotFound This error is returned when a product is not found
	ErrProductNotFound = errors.New("product not found")

	// ErrInvalidID This error is returned when the ID is invalid
	ErrInvalidID = errors.New("invalid ID")

	// ErrInvalidName This error is returned when the name is invalid
	ErrInvalidName = errors.New("invalid name")

	// ErrInvalidQuantity This error is returned when the quantity is invalid
	ErrInvalidQuantity = errors.New("invalid quantity regarding type, expected int")

	// ErrNegativeQuantity This error is returned when the quantity is negative
	ErrNegativeQuantity = errors.New("quantity cannot be negative")

	// ErrInvalidCodeValue This error is returned when the code value is invalid
	ErrInvalidCodeValue = errors.New("invalid code value")

	// ErrRepeatedCodeValue This error is returned when the code value already exists
	ErrRepeatedCodeValue = errors.New("code value already exists")

	// ErrInvalidIsPublished This error is returned when the is published is neither true nor false
	ErrInvalidIsPublished = errors.New("invalid is_published, expected true or false")

	// ErrInvalidExpiration This error is returned when the expiration is invalid
	ErrInvalidExpiration = errors.New("invalid expiration, expected date in format dd/mm/yyyy")

	// ErrInvalidPrice This error is returned when the price is invalid
	ErrInvalidPrice = errors.New("invalid price, expected decimal number")
)

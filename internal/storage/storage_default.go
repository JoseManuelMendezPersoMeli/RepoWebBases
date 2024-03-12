package storage

import (
	"time"
)

// Here we treat every function as if we were working with a desktop application

// ProductAttributesDefault Struct that has the default attributes of a product
type ProductAttributesDefault struct {
	Name        string
	Quantity    int
	CodeValue   string
	IsPublished bool
	Expiration  time.Time
	Price       float64
}

// StoredProductDefault Struct that has the information of a product on storage
type StoredProductDefault struct {
	warehouse    map[int]*ProductAttributesDefault
	availableIDs []int
}

// NewStoredProductDefault Creates a new StorageProductDefault+
func NewStoredProductDefault(warehouse map[int]*ProductAttributesDefault, availableIDs []int) *StoredProductDefault {
	return &StoredProductDefault{warehouse: warehouse, availableIDs: availableIDs}
}

func (storedProduct *StoredProductDefault) Ping() (string, error) {
	return "pong", nil
}

func (storedProduct *StoredProductDefault) GetAll() ([]*Product, error) {
	var products []*Product
	for id, product := range storedProduct.warehouse {
		products = append(products, &Product{
			ID:          id + 1,
			Name:        product.Name,
			Quantity:    product.Quantity,
			CodeValue:   product.CodeValue,
			IsPublished: product.IsPublished,
			Expiration:  product.Expiration,
			Price:       product.Price,
		})
	}
	return products, nil
}

func (storedProduct *StoredProductDefault) GetByID(id int) (*Product, error) {
	product, ok := storedProduct.warehouse[id-1]
	if !ok {
		return nil, ErrProductNotFound
	}
	return &Product{
		ID:          id,
		Name:        product.Name,
		Quantity:    product.Quantity,
		CodeValue:   product.CodeValue,
		IsPublished: product.IsPublished,
		Expiration:  product.Expiration,
		Price:       product.Price,
	}, nil
}

func (storedProduct *StoredProductDefault) Search(query *Query) ([]*Product, error) {
	products := storedProduct.warehouse
	if products == nil {
		return nil, ErrDatabaseAccess
	}

	var foundProducts []*Product
	for id, product := range products {
		if query != nil && query.PriceGt > 0 {
			if product.Price > query.PriceGt {
				foundProducts = append(foundProducts, &Product{id, product.Name, product.Quantity, product.CodeValue, product.IsPublished, product.Expiration, product.Price})
			}
		}
	}
	return foundProducts, nil
}

func (storedProduct *StoredProductDefault) AddProduct(product *ProductAttributesDefault) error {
	if product.Quantity < 0 {
		return ErrNegativeQuantity
	}
	if product.CodeValue != "" {
		for _, prod := range storedProduct.warehouse {
			if prod.CodeValue == product.CodeValue {
				return ErrRepeatedCodeValue
			}
		}
	}
	var ID int
	if len(storedProduct.availableIDs) == 0 {
		ID = len(storedProduct.warehouse)
	} else {
		ID = storedProduct.availableIDs[0]
		storedProduct.availableIDs = storedProduct.availableIDs[1:]
	}
	storedProduct.warehouse[ID] = product
	return nil
}

func (storedProduct *StoredProductDefault) UpdateOrCreateProduct(id int, product *ProductAttributesDefault) error {
	var ID int
	if id > len(storedProduct.warehouse) {
		ID = len(storedProduct.warehouse)
	} else {
		ID = id - 1
	}
	for _, prod := range storedProduct.warehouse {
		if prod.CodeValue == product.CodeValue {
			return ErrRepeatedCodeValue
		}
	}
	println(ID)
	storedProduct.warehouse[ID] = product
	return nil
}

func (storedProduct *StoredProductDefault) UpdateProduct(id int, product map[string]any) (*Product, error) {
	oldProduct, found := (*storedProduct).warehouse[id-1]
	if !found {
		return nil, ErrProductNotFound
	}
	for key, value := range product {
		switch key {
		case "Name", "name":
			name, ok := value.(string)
			if !ok {
				return nil, ErrInvalidName
			}
			oldProduct.Name = name
		case "Quantity", "quantity":
			quantity, ok := value.(int)
			if !ok {
				return nil, ErrInvalidQuantity
			}
			if quantity < 0 {
				return nil, ErrNegativeQuantity
			}
			oldProduct.Quantity = quantity
		case "CodeValue", "code_value":
			codeValue, ok := value.(string)
			if !ok {
				return nil, ErrInvalidCodeValue
			}
			for _, prod := range storedProduct.warehouse {
				if prod.CodeValue == codeValue {
					return nil, ErrRepeatedCodeValue
				}
			}
			oldProduct.CodeValue = codeValue
		case "IsPublished", "is_published":
			isPublished, ok := value.(bool)
			if !ok {
				return nil, ErrInvalidIsPublished
			}
			oldProduct.IsPublished = isPublished
		case "Expiration", "expiration":
			expiration, ok := value.(time.Time)
			if !ok {
				return nil, ErrInvalidExpiration
			}
			oldProduct.Expiration = expiration
		case "Price", "price":
			price, ok := value.(float64)
			if !ok {
				return nil, ErrInvalidPrice
			}
			oldProduct.Price = price
		}
	}
	newProduct := &Product{
		ID:          id,
		Name:        oldProduct.Name,
		Quantity:    oldProduct.Quantity,
		CodeValue:   oldProduct.CodeValue,
		IsPublished: oldProduct.IsPublished,
		Expiration:  oldProduct.Expiration,
		Price:       oldProduct.Price,
	}
	return newProduct, nil
}

func (storedProduct *StoredProductDefault) DeleteProduct(id int) (*Product, error) {
	product, found := (*storedProduct).warehouse[id-1]
	if !found {
		return nil, ErrProductNotFound
	}
	delete((*storedProduct).warehouse, id-1)
	(*storedProduct).availableIDs = append((*storedProduct).availableIDs, id-1)
	return &Product{
		ID:          id,
		Name:        product.Name,
		Quantity:    product.Quantity,
		CodeValue:   product.CodeValue,
		IsPublished: product.IsPublished,
		Expiration:  product.Expiration,
		Price:       product.Price,
	}, nil
}

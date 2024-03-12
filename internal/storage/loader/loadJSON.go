package loader

import (
	"RepoRefactor/internal/storage"
	"encoding/json"
	"os"
	"time"
)

func NewLoaderJSON(filePath string) *LoaderJSON {
	return &LoaderJSON{filePath: filePath}
}

type LoaderJSON struct {
	filePath string
}

type ProductsAttributesJSON struct {
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

type ProductJSON struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

func (l *LoaderJSON) Load() (*ProductsWarehouse, error) {
	// open file
	file, err := os.Open(l.filePath)
	if err != nil {
		return nil, ErrFileNotFound
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	// Products Aux
	var productsAux []ProductJSON

	warehouse := make(map[int]*storage.ProductAttributesDefault)
	err = json.NewDecoder(file).Decode(&productsAux)
	if err != nil {
		println(err.Error())
		return nil, ErrLoadingFile
	}

	dateLayout := "02/01/2006"

	for key, value := range productsAux {
		var expTime time.Time
		if value.Expiration != "" {
			expTime, err = time.Parse(dateLayout, value.Expiration)
			if err != nil {
				return nil, storage.ErrInvalidExpiration
			}
		}
		warehouse[key] = &storage.ProductAttributesDefault{
			Name:        value.Name,
			Quantity:    value.Quantity,
			CodeValue:   value.CodeValue,
			IsPublished: value.IsPublished,
			Expiration:  expTime,
			Price:       value.Price,
		}
	}

	var availableIDs []int

	wh := &ProductsWarehouse{
		Wh:           warehouse,
		AvailableIDs: availableIDs,
	}

	return wh, nil
}

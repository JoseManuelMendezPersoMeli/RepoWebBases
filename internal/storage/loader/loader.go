package loader

import (
	"RepoRefactor/internal/storage"
	"errors"
)

type ProductsWarehouse struct {
	Wh           map[int]*storage.ProductAttributesDefault
	AvailableIDs []int
}

type Loader interface {
	Load() (*ProductsWarehouse, error)
}

var (
	ErrFileNotFound = errors.New("file not found")
	ErrLoadingFile  = errors.New("error loading file")
)

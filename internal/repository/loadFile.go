package repository

import (
	"GoWeb/internal"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// LoadFile This function will load the JSON file to a slice and return it.
func LoadFile() (internal.ProductManager, error) {
	// Initialize the slice
	var productsManager internal.ProductManager
	var products []internal.Product

	// Bring the file from the docs folder
	sourceFolder := "docs"
	fileName := "products.json"

	sourcePath := filepath.Join(sourceFolder, fileName)

	// Read the file
	fileData, err := os.ReadFile(sourcePath)
	if err != nil {
		fmt.Println("Error reading file: ", err.Error())
		return internal.ProductManager{}, err
	}

	// Unmarshal the file data into the slice
	err = json.Unmarshal(fileData, &products)
	if err != nil {
		fmt.Println("Error unmarshaling file data: ", err)
		return internal.ProductManager{}, err
	}
	productsManager.Products = products

	// Return the slice and nil error
	return productsManager, nil
}

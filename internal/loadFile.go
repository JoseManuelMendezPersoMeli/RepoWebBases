package internal

import (
	"encoding/json"
	"fmt"
	"os"
)

// LoadFile This function will load the JSON file to a slice and return it.
func LoadFile() ([]Product, error) {
	// Initialize the slice
	var products []Product

	// Read the file
	fileData, err := os.ReadFile("products.json")
	if err != nil {
		fmt.Println("Error reading file: ", err.Error())
		return nil, err
	}

	// Unmarshal the file data into the slice
	err = json.Unmarshal(fileData, &products)
	if err != nil {
		fmt.Println("Error unmarshaling file data: ", err)
		return nil, err
	}

	// Return the slice and nil error
	return products, nil
}

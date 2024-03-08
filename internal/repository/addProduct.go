package repository

import (
	"GoWeb/internal"
	"GoWeb/internal/service"
	"GoWeb/platform/request"
	"GoWeb/platform/response"
	"net/http"
)

func AddProduct(products *internal.ProductManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Request
		var body internal.ProductRequestBody
		var nextID int

		// - Read the request body
		if err := request.RequestJSON(r, &body); err != nil {
			response.ResponseJSON(w, http.StatusBadRequest, map[string]any{"error": "Invalid request body"})
			return
		}

		// - Validate required keys
		err := service.ValidateKeys(body)
		if err != nil {
			response.ResponseJSON(w, http.StatusBadRequest, map[string]any{"error": err.Error()})
			return
		}

		// - Parse to map (static)

		// Process
		// - Get next ID
		if len(products.AvailableIDs) == 0 {
			nextID = len(products.Products) + 1
		} else {
			nextID = products.AvailableIDs[0]
			products.AvailableIDs = products.AvailableIDs[1:]
		}

		// - Serialize the request body into a Product struct
		product := internal.Product{
			ID:          nextID,
			Name:        body.Name,
			Quantity:    body.Quantity,
			CodeValue:   body.CodeValue,
			IsPublished: body.IsPublished,
			Expiration:  body.Expiration,
			Price:       body.Price,
		}

		productJSON := internal.ProductJSON{
			ID:          product.ID,
			Name:        product.Name,
			Quantity:    product.Quantity,
			CodeValue:   product.CodeValue,
			IsPublished: product.IsPublished,
			Expiration:  product.Expiration,
			Price:       product.Price,
		}

		// - Add the product to the product manager
		products.Products = append(products.Products, product)

		// Response
		response.ResponseJSON(w, http.StatusCreated, map[string]any{
			"message": "Product added successfully",
			"data":    productJSON,
		})
		println(products.Products)
	}
}

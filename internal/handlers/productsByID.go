package handlers

import (
	"GoWeb/internal"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func ProductsByIDHandler(products []internal.Product) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := strings.TrimPrefix(r.URL.Path, "/products/")
		id := strings.Trim(path, "/")

		IDInt, _ := strconv.Atoi(id)
		var foundProduct *internal.Product
		for _, product := range products {
			if product.ID == IDInt {
				foundProduct = &product
				break
			}
		}

		if foundProduct != nil {
			w.Header().Set("Content-Type", "application/json")
			err := json.NewEncoder(w).Encode(foundProduct)
			if err != nil {
				http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
				return
			}
		} else {
			http.Error(w, "Product not found", http.StatusNotFound)
			return
		}
	}
}

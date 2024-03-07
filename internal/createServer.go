package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func PingHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprint(w, "Pong!", http.StatusOK)
	if err != nil {
		http.Error(w, "Error writing response", http.StatusInternalServerError)
		return
	}

}

func ProductsHandler(products []Product) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		err := json.NewEncoder(w).Encode(products)
		if err != nil {
			http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
			return
		}
	}
}

func ProductsByIDHandler(products []Product) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := strings.TrimPrefix(r.URL.Path, "/products/")
		id := strings.Trim(path, "/")

		IDInt, _ := strconv.Atoi(id)
		var foundProduct *Product
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

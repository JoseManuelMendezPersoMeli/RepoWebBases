package handlers

import (
	"GoWeb/internal"
	"encoding/json"
	"net/http"
	"strconv"
)

func SearchHandler(products *internal.ProductManager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		priceGt := r.URL.Query().Get("priceGt")
		if priceGt != "" {
			priceGtInt, _ := strconv.Atoi(priceGt)
			var foundProducts []internal.Product
			for _, product := range products.Products {
				if product.Price > float64(priceGtInt) {
					foundProducts = append(foundProducts, product)
				}
			}
			w.Header().Set("Content-Type", "application/json")
			err := json.NewEncoder(w).Encode(foundProducts)
			if err != nil {
				http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
				return
			}
		} else {
			http.Error(w, "Invalid price", http.StatusBadRequest)
			return
		}
	}
}

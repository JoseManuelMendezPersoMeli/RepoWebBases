package handlers

import (
	"GoWeb/internal"
	"encoding/json"
	"net/http"
)

func ProductsHandler(products []internal.Product) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		err := json.NewEncoder(w).Encode(products)
		if err != nil {
			http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
			return
		}
	}
}

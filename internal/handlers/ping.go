package handlers

import (
	"fmt"
	"net/http"
)

func PingHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprint(w, "Pong!", http.StatusOK)
	if err != nil {
		http.Error(w, "Error writing response", http.StatusInternalServerError)
		return
	}

}

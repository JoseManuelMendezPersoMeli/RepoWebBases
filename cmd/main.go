package main

import (
	"GoWeb/internal"
	"GoWeb/internal/handlers"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {

	// Dependency injection
	file, err := internal.LoadFile()
	if err != nil {
		println("Error loading file: ", err.Error())
		return
	}

	// Create a new router
	router := chi.NewRouter()
	router.Get("/ping", handlers.PingHandler)
	router.Get("/products", handlers.ProductsHandler(file))
	router.Get("/products/{id:[0-9]+}", handlers.ProductsByIDHandler(file))
	router.Get("/products/search", handlers.SearchHandler(file))

	http.Handle("/", router)

	println("Server started on port 8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		println("Error starting server: ", err.Error())
		return
	}

}

package main

import (
	"GoWeb/internal/handlers"
	"GoWeb/internal/repository"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {

	// Dependency injection
	productManager, err := repository.LoadFile()
	if err != nil {
		println("Error loading file: ", err.Error())
		return
	}

	println(&productManager.Products)

	// Create a new router
	router := chi.NewRouter()
	router.Get("/ping", handlers.PingHandler)
	router.Get("/products", handlers.ProductsHandler(&productManager))
	router.Get("/products/{id:[0-9]+}", handlers.ProductsByIDHandler(&productManager))
	router.Get("/products/search", handlers.SearchHandler(&productManager))
	router.Post("/products", repository.AddProduct(&productManager))

	http.Handle("/", router)

	println("Server started on port 8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		println("Error starting server: ", err.Error())
		return
	}

}

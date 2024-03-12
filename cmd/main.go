package main

import (
	"RepoRefactor/cmd/handlers"
	"RepoRefactor/internal/storage"
	"RepoRefactor/internal/storage/loader"
	"github.com/go-chi/chi/v5"
	"net/http"
	"path/filepath"
)

func main() {
	// Dependencies
	sourceFolder := "docs"
	fileName := "products.json"
	sourcePath := filepath.Join(sourceFolder, fileName)

	ld := loader.NewLoaderJSON(sourcePath)
	dataBase, err := ld.Load()
	if err != nil {
		println("Error loading file: ", err.Error())
		return
	}
	store := storage.NewStoredProductDefault(dataBase.Wh, dataBase.AvailableIDs)
	controller := handlers.NewProductsController(store)

	// server
	router := chi.NewRouter()
	// Handlers
	router.Get("/ping", controller.Ping())
	router.Route("/products", func(router chi.Router) {
		router.Get("/", controller.GetAll())
		router.Get("/{id:[0-9]+}", controller.GetByID())
		router.Get("/search", controller.Search())
		router.Post("/", controller.AddProduct())
		router.Put("/{id:[0-9]+}", controller.UpdateOrCreateProduct())
		router.Patch("/{id:[0-9]+}", controller.UpdateProduct())
		router.Delete("/{id:[0-9]+}", controller.DeleteProduct())
	})

	if err := http.ListenAndServe(":8080", router); err != nil {
		println("Error starting server: ", err.Error())
		return
	}
}

package main

import (
	"GoWeb/internal"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	file, err := internal.LoadFile()
	if err != nil {
		println("Error loading file: ", err.Error())
		return
	}

	r := mux.NewRouter()
	r.HandleFunc("/ping", internal.PingHandler)
	r.HandleFunc("/products", internal.ProductsHandler(file))
	r.HandleFunc("/products/{id:[0-9]+}", internal.ProductsByIDHandler(file))

	http.Handle("/", r)

	println("Server started on port 8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		println("Error starting server: ", err.Error())
		return
	}

}

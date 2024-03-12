package handlers

import (
	"RepoRefactor/internal/storage"
	"RepoRefactor/internal/validator"
	"RepoRefactor/platform/web/request"
	"RepoRefactor/platform/web/response"
	"errors"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
	"time"
)

func NewProductsController(st storage.StoredProduct) *ProductsController {
	return &ProductsController{st: st}
}

type ProductsController struct {
	st storage.StoredProduct
}

type ResponsePong struct {
	Message string `json:"message"`
	Error   bool   `json:"error"`
}

func (c *ProductsController) Ping() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pong, err := c.st.Ping()
		if err != nil {
			code := http.StatusInternalServerError
			body := &ResponsePong{Message: "Error while ping", Error: false}
			response.ResponseJSON(w, code, body)
		}
		code := http.StatusOK
		body := &ResponsePong{Message: pong, Error: false}
		response.ResponseJSON(w, code, body)
	}
}

type ProductHandlerGetAll struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

type ResponseGetAll struct {
	Message string                  `json:"message"`
	Data    []*ProductHandlerGetAll `json:"data"`
	Error   bool                    `json:"error"`
}

func (c *ProductsController) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		products, err := c.st.GetAll()
		if err != nil {
			code := http.StatusInternalServerError
			body := &ResponseGetAll{Message: err.Error(), Data: nil, Error: true}

			response.ResponseJSON(w, code, body)
			return
		}

		code := http.StatusOK
		body := &ResponseGetAll{Message: "Success", Data: make([]*ProductHandlerGetAll, len(products)), Error: false}
		for i, product := range products {
			body.Data[i] = &ProductHandlerGetAll{
				ID:          product.ID,
				Name:        product.Name,
				Quantity:    product.Quantity,
				CodeValue:   product.CodeValue,
				IsPublished: product.IsPublished,
				Expiration:  product.Expiration.Format("29/01/2006"),
				Price:       product.Price,
			}
		}
		response.ResponseJSON(w, code, body)
	}
}

type ProductHandlerGetByID struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

type ResponseGetByID struct {
	Message string                 `json:"message"`
	Data    *ProductHandlerGetByID `json:"data"`
	Error   bool                   `json:"error"`
}

func (c *ProductsController) GetByID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Request
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			code := http.StatusBadRequest
			body := &ResponseGetByID{Message: err.Error(), Data: nil, Error: true}
			response.ResponseJSON(w, code, body)
			return
		}

		// Process
		product, err := c.st.GetByID(id)
		if err != nil {
			var code int
			var body *ResponseGetByID
			switch {
			case errors.Is(err, storage.ErrProductNotFound):
				code = http.StatusNotFound
				body = &ResponseGetByID{Message: err.Error(), Data: nil, Error: true}
			default:
				code = http.StatusInternalServerError
				body = &ResponseGetByID{Message: err.Error(), Data: nil, Error: true}
			}
			response.ResponseJSON(w, code, body)
			return
		}

		// Response
		code := http.StatusOK
		body := &ResponseGetByID{Message: "Success", Data: &ProductHandlerGetByID{
			ID:          product.ID,
			Name:        product.Name,
			Quantity:    product.Quantity,
			CodeValue:   product.CodeValue,
			IsPublished: product.IsPublished,
			Expiration:  product.Expiration.Format("29/01/2006"),
			Price:       product.Price,
		}, Error: false}

		response.ResponseJSON(w, code, body)
	}
}

type ProductHandlerSearch struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

type ResponseSearch struct {
	Message string                  `json:"message"`
	Data    []*ProductHandlerSearch `json:"data"`
	Error   bool                    `json:"error"`
}

type ProductQueryHandlerSearch struct {
	PriceGt float64 `json:"price_gt"`
}

func (c *ProductsController) Search() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Request
		var query ProductQueryHandlerSearch
		query.PriceGt, _ = strconv.ParseFloat(r.URL.Query().Get("price_gt"), 64)

		// Process
		q := &storage.Query{PriceGt: query.PriceGt}
		products, err := c.st.Search(q)
		if err != nil {
			code := http.StatusInternalServerError
			body := &ResponseSearch{Message: err.Error(), Data: nil, Error: true}
			response.ResponseJSON(w, code, body)
			return
		}

		// Response
		code := http.StatusOK
		body := &ResponseSearch{Message: "Success", Data: make([]*ProductHandlerSearch, len(products)), Error: false}
		for i, product := range products {
			body.Data[i] = &ProductHandlerSearch{
				ID:          product.ID,
				Name:        product.Name,
				Quantity:    product.Quantity,
				CodeValue:   product.CodeValue,
				IsPublished: product.IsPublished,
				Expiration:  product.Expiration.Format("29/01/2006"),
				Price:       product.Price,
			}
		}
		response.ResponseJSON(w, code, body)
	}
}

type ProductHandlerAddProduct struct {
	Name        string  `json:"name"`
	Quantity    int     `json:"quantity"`
	CodeValue   string  `json:"code_value"`
	IsPublished bool    `json:"is_published"`
	Expiration  string  `json:"expiration"`
	Price       float64 `json:"price"`
}

type ResponseAddProduct struct {
	Message string                    `json:"message"`
	Data    *ProductHandlerAddProduct `json:"data"`
	Error   bool                      `json:"error"`
}

func (c *ProductsController) AddProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Request
		var productBody ProductHandlerAddProduct
		err := request.RequestJSON(r, &productBody)
		if err != nil {
			code := http.StatusBadRequest
			body := &ResponseAddProduct{Message: err.Error(), Data: nil, Error: true}
			response.ResponseJSON(w, code, body)
			return
		}

		// Process

		// Save
		dateLayout := "02/01/2006"
		exp, err := time.Parse(dateLayout, productBody.Expiration)
		if err != nil {
			code := http.StatusBadRequest
			body := &ResponseAddProduct{Message: "Invalid date format", Data: nil, Error: true}
			response.ResponseJSON(w, code, body)
		}
		product := &storage.ProductAttributesDefault{
			Name:        productBody.Name,
			Quantity:    productBody.Quantity,
			CodeValue:   productBody.CodeValue,
			IsPublished: productBody.IsPublished,
			Expiration:  exp,
			Price:       productBody.Price,
		}
		err = c.st.AddProduct(product)
		if err != nil {
			code := http.StatusInternalServerError
			body := &ResponseAddProduct{Message: err.Error(), Data: nil, Error: true}
			response.ResponseJSON(w, code, body)
			return
		}

		// Response
		// - Change types to show the data
		productAdded := &ProductHandlerAddProduct{
			Name:        product.Name,
			Quantity:    product.Quantity,
			CodeValue:   product.CodeValue,
			IsPublished: product.IsPublished,
			Expiration:  product.Expiration.Format("02/01/2006"),
			Price:       product.Price,
		}
		code := http.StatusOK
		body := &ResponseAddProduct{Message: "Success", Data: productAdded, Error: false}
		response.ResponseJSON(w, code, body)
	}
}

type ProductHandlerUpdateOrCreateProduct struct {
	Name        string  `json:"name,omitempty"`
	Quantity    int     `json:"quantity,omitempty"`
	CodeValue   string  `json:"code_value,omitempty"`
	IsPublished bool    `json:"is_published,omitempty"`
	Expiration  string  `json:"expiration,omitempty"`
	Price       float64 `json:"price,omitempty"`
}

type ResponseUpdateOrCreateProduct struct {
	Message string                    `json:"message"`
	Data    *ProductHandlerAddProduct `json:"data"`
	Error   bool                      `json:"error"`
}

type ProductQueryHandlerUpdateOrCreate struct {
	ID int `json:"id"`
}

func (c *ProductsController) UpdateOrCreateProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Request
		var query ProductQueryHandlerUpdateOrCreate
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			code := http.StatusBadRequest
			body := &ResponseUpdateOrCreateProduct{Message: err.Error(), Data: nil, Error: true}
			response.ResponseJSON(w, code, body)
			return
		}
		query.ID = id

		var productBody ProductHandlerUpdateOrCreateProduct
		err = request.RequestJSON(r, &productBody)
		if err != nil {
			code := http.StatusBadRequest
			body := &ResponseUpdateOrCreateProduct{Message: err.Error(), Data: nil, Error: true}
			response.ResponseJSON(w, code, body)
			return
		}

		// Process
		// Save
		dateLayout := "02/01/2006"
		exp, err := time.Parse(dateLayout, productBody.Expiration)
		if err != nil {
			code := http.StatusBadRequest
			body := &ResponseUpdateOrCreateProduct{Message: "Invalid date format, must be dd/mm/yyyy", Data: nil, Error: true}
			response.ResponseJSON(w, code, body)
			return
		}
		product := &storage.ProductAttributesDefault{
			Name:        productBody.Name,
			Quantity:    productBody.Quantity,
			CodeValue:   productBody.CodeValue,
			IsPublished: productBody.IsPublished,
			Expiration:  exp,
			Price:       productBody.Price,
		}

		productValidated := validator.ValidatorProductDetails{
			Name:        product.Name,
			Quantity:    product.Quantity,
			CodeValue:   product.CodeValue,
			IsPublished: product.IsPublished,
			Expiration:  product.Expiration,
			Price:       product.Price,
		}
		err = validator.ValidateProductDetails(&productValidated)
		if err != nil {
			code := http.StatusBadRequest
			body := &ResponseUpdateOrCreateProduct{Message: err.Error(), Data: nil, Error: true}
			response.ResponseJSON(w, code, body)
			return
		}
		err = c.st.UpdateOrCreateProduct(query.ID, product)
		if err != nil {
			code := http.StatusInternalServerError
			body := &ResponseAddProduct{Message: err.Error(), Data: nil, Error: true}
			response.ResponseJSON(w, code, body)
			return
		}

		// Response
		// - Change types to show the data
		productAdded := &ProductHandlerAddProduct{
			Name:        product.Name,
			Quantity:    product.Quantity,
			CodeValue:   product.CodeValue,
			IsPublished: product.IsPublished,
			Expiration:  product.Expiration.Format("02/01/2006"),
			Price:       product.Price,
		}
		code := http.StatusOK
		body := &ResponseAddProduct{Message: "Success", Data: productAdded, Error: false}
		response.ResponseJSON(w, code, body)
	}
}

type ProductHandlerUpdateProduct struct {
	Name        string  `json:"name,omitempty"`
	Quantity    int     `json:"quantity,omitempty"`
	CodeValue   string  `json:"code_value,omitempty"`
	IsPublished bool    `json:"is_published,omitempty"`
	Expiration  string  `json:"expiration,omitempty"`
	Price       float64 `json:"price"`
}

type ResponseUpdateProduct struct {
	Message string                       `json:"message"`
	Data    *ProductHandlerUpdateProduct `json:"data"`
	Error   bool                         `json:"error"`
}

type ProductQueryHandlerUpdate struct {
	ID int `json:"id"`
}

func (c *ProductsController) UpdateProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Request
		var query ProductQueryHandlerUpdate
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			code := http.StatusBadRequest
			body := &ResponseUpdateProduct{Message: err.Error(), Data: nil, Error: true}
			response.ResponseJSON(w, code, body)
			return
		}
		query.ID = id

		productBody := make(map[string]any)
		err = request.RequestJSON(r, &productBody)
		if err != nil {
			code := http.StatusBadRequest
			body := &ResponseUpdateProduct{Message: err.Error(), Data: nil, Error: true}
			response.ResponseJSON(w, code, body)
			return
		}

		// Process
		// Save
		newProduct, err := c.st.UpdateProduct(query.ID, productBody)
		if err != nil {
			code := http.StatusInternalServerError
			body := &ResponseUpdateProduct{Message: err.Error(), Data: nil, Error: true}
			response.ResponseJSON(w, code, body)
			return
		}
		updatedProduct := &ProductHandlerUpdateProduct{
			Name:        newProduct.Name,
			Quantity:    newProduct.Quantity,
			CodeValue:   newProduct.CodeValue,
			IsPublished: newProduct.IsPublished,
			Expiration:  newProduct.Expiration.Format("02/01/2006"),
			Price:       newProduct.Price,
		}
		code := http.StatusOK
		body := &ResponseUpdateProduct{Message: "Success", Data: updatedProduct, Error: false}
		response.ResponseJSON(w, code, body)
	}
}

type ProductHandlerDeleteProduct struct {
	Name        string  `json:"name,omitempty"`
	Quantity    int     `json:"quantity,omitempty"`
	CodeValue   string  `json:"code_value,omitempty"`
	IsPublished bool    `json:"is_published,omitempty"`
	Expiration  string  `json:"expiration,omitempty"`
	Price       float64 `json:"price"`
}

type ResponseDeleteProduct struct {
	Message string                       `json:"message"`
	Data    *ProductHandlerDeleteProduct `json:"data"`
	Error   bool                         `json:"error"`
}

type ProductQueryHandlerDelete struct {
	ID int `json:"id"`
}

func (c *ProductsController) DeleteProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Request
		var query ProductQueryHandlerDelete
		id, err := strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			code := http.StatusBadRequest
			body := &ResponseDeleteProduct{Message: err.Error(), Data: nil, Error: true}
			response.ResponseJSON(w, code, body)
			return
		}
		query.ID = id

		// Process
		// Save
		product, err := c.st.DeleteProduct(query.ID)
		if err != nil {
			code := http.StatusInternalServerError
			body := &ResponseDeleteProduct{Message: err.Error(), Data: nil, Error: true}
			response.ResponseJSON(w, code, body)
			return
		}
		deletedProduct := &ProductHandlerDeleteProduct{
			Name:        product.Name,
			Quantity:    product.Quantity,
			CodeValue:   product.CodeValue,
			IsPublished: product.IsPublished,
			Expiration:  product.Expiration.Format("02/01/2006"),
			Price:       product.Price,
		}
		code := http.StatusOK
		body := &ResponseDeleteProduct{Message: "Success", Data: deletedProduct, Error: false}
		response.ResponseJSON(w, code, body)
	}
}

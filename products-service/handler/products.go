package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/subhankardas/go-microservices/products-service/data"
)

type Products struct {
	log *log.Logger
}

// Create a new products handler
func NewProductsHandler(log *log.Logger) *Products {
	return &Products{log}
}

// Method to get list of products and write to response as json
func (prd *Products) GetProducts(response http.ResponseWriter, request *http.Request) {
	prd.log.Print("Handle GET request for products.")
	products := data.GetProducts()

	// Parse data to json and write to response
	err := products.ToJSON(response)
	if err != nil {
		prd.log.Print("Unable to parse products data.")
		http.Error(response, "Invalid products data.", http.StatusInternalServerError)
	}
}

// Method to add product to list
func (prd *Products) AddProduct(response http.ResponseWriter, request *http.Request) {
	prd.log.Print("Handle POST request for product.")
	product := &data.Product{}

	// Parse json data from request
	err := product.FromJSON(request.Body)
	if err != nil {
		prd.log.Print("Unable to parse product data.")
		http.Error(response, "Invalid product data.", http.StatusBadRequest)
		return
	}

	// Add product data
	data.AddProduct(product)
}

// Method to update product data
func (prd *Products) UpdateProduct(response http.ResponseWriter, request *http.Request) {
	prd.log.Print("Handle PUT request for product.")

	// Get request variables
	vars := mux.Vars(request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		prd.log.Print("Invalid product ID passed for update.")
		http.Error(response, "Invalid product ID.", http.StatusBadRequest)
		return
	}

	// Parse json data from request
	product := &data.Product{}
	err = product.FromJSON(request.Body)
	if err != nil {
		prd.log.Print("Unable to parse product data.")
		http.Error(response, "Invalid product data.", http.StatusBadRequest)
		return
	}

	// Update product data
	err = data.UpdateProduct(id, product)
	if err != nil {
		prd.log.Printf("[ERROR] %v!", err)
		http.Error(response, "Unable to update, product not found.", http.StatusNotFound)
		return
	}
}

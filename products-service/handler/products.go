package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/subhankardas/go-microservices/products-service/data"
)

type Products struct {
	log *log.Logger
}

// Create a new products handler
func NewProductsHandler(log *log.Logger) *Products {
	return &Products{log}
}

// Implement the ServeHTTP() method by Handler interface to make this struct a handler
func (prd *Products) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	// Handle GET request and return list of products
	if request.Method == http.MethodGet {
		prd.getProducts(response, request)
		return
	}

	// Handle all the http request and return not allowed header
	response.WriteHeader(http.StatusMethodNotAllowed)
}

// Method to get list of products and write to response as json
func (prd *Products) getProducts(response http.ResponseWriter, request *http.Request) {
	products := data.GetProducts()

	// Parse data to json
	data, err := json.Marshal(products)
	if err != nil {
		prd.log.Print("Unable to convert data, invalid format.")
		http.Error(response, "Invalid product data.", http.StatusInternalServerError)
	}

	// Write json data to response
	response.Write(data)
}

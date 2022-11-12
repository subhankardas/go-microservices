package handler

import (
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
		prd.log.Print("Handle GET request for products.")
		prd.getProducts(response, request)
		return
	} else if request.Method == http.MethodPost {
		prd.log.Print("Handle POST request for product.")
		prd.addProduct(response, request)
		return
	} else if request.Method == http.MethodPut {
		prd.log.Print("Handle PUT request for product.")
		prd.updateProduct(response, request)
		return
	}

	// Handle all other http requests and return not allowed header
	response.WriteHeader(http.StatusMethodNotAllowed)
}

// Method to get list of products and write to response as json
func (prd *Products) getProducts(response http.ResponseWriter, request *http.Request) {
	products := data.GetProducts()

	// Parse data to json and write to response
	err := products.ToJSON(response)
	if err != nil {
		prd.log.Print("Unable to parse products data.")
		http.Error(response, "Invalid products data.", http.StatusInternalServerError)
	}
}

// Method to add product to list
func (prd *Products) addProduct(response http.ResponseWriter, request *http.Request) {
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
func (prd *Products) updateProduct(response http.ResponseWriter, request *http.Request) {
	product := &data.Product{}

	// Parse json data from request
	err := product.FromJSON(request.Body)
	if err != nil {
		prd.log.Print("Unable to parse product data.")
		http.Error(response, "Invalid product data.", http.StatusBadRequest)
		return
	}

	// Update product data
	err = data.UpdateProduct(product)
	if err != nil {
		prd.log.Printf("[ERROR] %v!", err)
		http.Error(response, "Unable to update, product not found.", http.StatusNotFound)
		return
	}
}

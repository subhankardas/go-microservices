package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/subhankardas/go-microservices/products-service/data"
	"github.com/subhankardas/go-microservices/products-service/middleware"
)

type Products struct {
	log *log.Logger
}

// Create a new products handler
func NewProductsHandler(log *log.Logger) *Products {
	return &Products{log}
}

// Method to get list of products and write to response as json.
// @Summary      Get list of products.
// @Description  Returns list of all the products in store.
// @Tags         products
// @Accept       json
// @Produce      json
// @Success      200  {object}  data.Products
// @Failure      500  {string}  string
// @Router       /products [get]
func (prd *Products) GetProducts(response http.ResponseWriter, request *http.Request) {
	prd.log.Print("Handle GET request for products.")
	response.Header().Add("Content-Type", "application/json")

	products := data.GetProducts()

	// Parse data to json and write to response
	err := products.ToJSON(response)
	if err != nil {
		prd.log.Print("Unable to parse products data.")
		http.Error(response, "Invalid products data.", http.StatusInternalServerError)
	}
}

// Method to add new product to list.
// @Summary 	Add new product.
// @Description Adds a new product to the store.
// @Tags        products
// @Accept     	json
// @Produce     json
// @Param product body data.Product true "Product Details"
// @Success     200  {object}  string
// @Failure     400  {object}  string
// @Router 		/products [post]
func (prd *Products) AddProduct(response http.ResponseWriter, request *http.Request) {
	prd.log.Print("Handle POST request for product.")

	// Get data from request and map to product
	product := request.Context().Value(middleware.KeyProduct{}).(*data.Product)

	// Add product data
	data.AddProduct(product)
}

// Method to update product data.
// @Summary 	Update existing product.
// @Description Update product details to the store.
// @Tags        products
// @Accept     	json
// @Produce     json
// @Param       id	path	int  true	"Product ID"
// @Param product body data.Product true "Product Details"
// @Success     200  {object}  string
// @Failure     400  {object}  string
// @Failure     404  {object}  string
// @Router 		/products/{id} [put]
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

	// Get data from request and map to product
	product := request.Context().Value(middleware.KeyProduct{}).(*data.Product)

	// Update product data
	err = data.UpdateProduct(id, product)
	if err != nil {
		prd.log.Printf("[ERROR] %v!", err)
		http.Error(response, "Unable to update, product not found.", http.StatusNotFound)
		return
	}
}

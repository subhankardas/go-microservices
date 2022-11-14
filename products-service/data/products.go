package data

import (
	"encoding/json"
	"errors"
	"io"
	"regexp"
	"time"

	"github.com/go-playground/validator/v10"
)

var (
	ErrorProductNotFound = errors.New("Product not found")
)

type Product struct {
	ID          int       `json:"id"`
	Name        string    `json:"name" validate:"required"`
	Description string    `json:"description"`
	Price       float32   `json:"price" validate:"gt=0"`
	SKU         string    `json:"sku" validate:"required,sku"` // SKU has custom validator
	CreatedOn   time.Time `json:"-"`                           // Ignore these fields when marshalling to json
	UpdatedOn   time.Time `json:"-"`
	DeletedOn   time.Time `json:"-"`
}

type Products []*Product

// Validates data constraints
func (prd *Product) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("sku", validateSKU) // Custom validation for SKU
	return validate.Struct(prd)
}

// Custom validator function to validate SKU
func validateSKU(field validator.FieldLevel) bool {
	value := field.Field().String()
	matched, err := regexp.MatchString("COFF[0-9]+", value) // SKU format must be COFF[0-9]+
	if err != nil {
		return false
	}
	return matched
}

// Reads json data from request
func (prd *Product) FromJSON(reader io.Reader) error {
	decoder := json.NewDecoder(reader)
	return decoder.Decode(prd)
}

// Writes json data to response
func (prd *Products) ToJSON(writer io.Writer) error {
	encoder := json.NewEncoder(writer)
	return encoder.Encode(prd)
}

// Data access to get products list
func GetProducts() Products {
	return products
}

// Data access to add product
func AddProduct(product *Product) {
	product.ID = len(products) + 1
	product.CreatedOn = time.Now().UTC()
	product.UpdatedOn = time.Time{}
	product.DeletedOn = time.Time{}

	products = append(products, product)
}

// Data access to update product
func UpdateProduct(id int, product *Product) error {
	var idx int = -1
	for i, prd := range products {
		if prd.ID == id {
			idx = i
			break
		}
	}
	if idx < 0 {
		return ErrorProductNotFound
	}

	// Update product details
	product.ID = id
	product.CreatedOn = products[idx].CreatedOn
	product.UpdatedOn = time.Now().UTC()
	products[idx] = product
	return nil
}

var products = []*Product{
	{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee.",
		Price:       3.45,
		SKU:         "COFF123",
		CreatedOn:   time.Now().UTC(),
		UpdatedOn:   time.Now().UTC(),
	},
	{
		ID:          2,
		Name:        "Espresso",
		Description: "Strong coffee without milk.",
		Price:       2.34,
		SKU:         "COFF456",
		CreatedOn:   time.Now().UTC(),
		UpdatedOn:   time.Now().UTC(),
	},
}

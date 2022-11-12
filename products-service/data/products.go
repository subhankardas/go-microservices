package data

import (
	"encoding/json"
	"io"
	"time"
)

type Product struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       float32   `json:"price"`
	SKU         string    `json:"sku"`
	CreatedOn   time.Time `json:"-"` // Ignore these fields when marshalling to json
	UpdatedOn   time.Time `json:"-"`
	DeletedOn   time.Time `json:"-"`
}

type Products []*Product

// Writes json data for this struct
func (prd *Products) ToJSON(writer io.Writer) error {
	encoder := json.NewEncoder(writer)
	return encoder.Encode(prd)
}

// Data access to products list
func GetProducts() Products {
	return products
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

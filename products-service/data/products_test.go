package data

import "testing"

func TestProductValidation(t *testing.T) {
	t.Run("test_product_fields_data_constraints", func(t *testing.T) {
		// Given
		product := &Product{
			Name:  "Latte",
			Price: 2.3,
			SKU:   "COFF123",
		}

		// When
		err := product.Validate()

		// Then
		if err != nil {
			t.Fatal(err)
		}
	})
}

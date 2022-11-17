package client

import (
	"testing"

	"github.com/subhankardas/go-microservices/products-service/client/products"
)

const (
	HOST = "localhost:8080"
)

func TestProductsClient(t *testing.T) {
	t.Run("test_products_get_api_client", func(t *testing.T) {
		// GIVEN - Create client with config
		config := DefaultTransportConfig().WithHost(HOST)
		client := NewHTTPClientWithConfig(nil, config)

		// WHEN - Call get products api via client
		params := products.NewGetProductsParams()

		// Then - Check error and response
		response, err := client.Products.GetProducts(params)
		products := response.GetPayload()

		if err != nil {
			t.Errorf("[ERROR] GET products call failed with error %v", err)
		}
		if len(products) == 0 {
			t.Error("[ERROR] GET products returned nil/empty response")
		}
	})
}

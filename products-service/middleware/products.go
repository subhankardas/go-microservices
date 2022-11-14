package middleware

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/subhankardas/go-microservices/products-service/data"
)

type Middleware struct {
	log *log.Logger
}

func New(log *log.Logger) *Middleware {
	return &Middleware{log}
}

type KeyProduct struct{}

func (mid *Middleware) ProductsMW(next func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(response http.ResponseWriter, req *http.Request) {
		// Parse json data from request
		product := &data.Product{}
		err := product.FromJSON(req.Body)
		if err != nil {
			mid.log.Print("Unable to parse product data.")
			http.Error(response, "Invalid product data.", http.StatusBadRequest)
			return
		}

		// Validate product details
		err = product.Validate()
		if err != nil {
			mid.log.Print("[ERROR] Product validation failed.")
			http.Error(response, fmt.Sprintf("Invalid product data. Error: %v", err.Error()), http.StatusBadRequest)
			return
		}

		// Inject data to request context
		ctx := req.Context()
		ctx = context.WithValue(ctx, KeyProduct{}, product)
		request := req.WithContext(ctx)

		next(response, request)
	}
}

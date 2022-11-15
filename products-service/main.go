// @title           Products Service API
// @version         1.0.0
// @description     API documentation for the products service.

// @contact.name   Subhankar Das
// @contact.url    https://github.com/subhankardas
// @contact.email  subhankardas831@gmail.com

// @host      localhost:8080
// @BasePath  /api
package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/subhankardas/go-microservices/products-service/handler"
	"github.com/subhankardas/go-microservices/products-service/middleware"
)

const (
	SERVER_ADDR = ":8080"
)

func main() {
	// Create required dependencies
	log := log.New(os.Stdout, "[products-service] ", log.LstdFlags)

	// Create handlers
	productsHandler := handler.NewProductsHandler(log)

	// Create new middleware
	middleware := middleware.New(log)

	// Setup handlers with custom router
	mux := mux.NewRouter()
	mux.HandleFunc("/api/products", productsHandler.GetProducts).Methods(http.MethodGet)
	mux.HandleFunc("/api/products", middleware.ProductsMW(productsHandler.AddProduct)).Methods(http.MethodPost)
	mux.HandleFunc("/api/products/{id}", middleware.ProductsMW(productsHandler.UpdateProduct)).Methods(http.MethodPut)

	// Create custom server and setup configs
	server := &http.Server{
		Addr:         SERVER_ADDR, // Server host and port
		Handler:      mux,         // Custom router mux
		IdleTimeout:  10 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		log.Printf("Starting server at %v.", SERVER_ADDR)

		// Start sever, listen to specified address and handle shutdown
		err := server.ListenAndServe()
		if err != nil {
			log.Print("Server shutdown successfully.")
		}
	}()

	// Make channel to listen to interrupt signal
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt)

	// Once interrupt signal is received proceed with graceful server shutdown
	shutdownSignal := <-signalCh
	log.Printf("Shutting down server gracefully, received signal %v.", shutdownSignal)

	// Shutdown server with timeout context
	timeoutCtx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	server.Shutdown(timeoutCtx)
}

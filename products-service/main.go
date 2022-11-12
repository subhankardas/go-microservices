package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/subhankardas/go-microservices/products-service/handler"
)

const (
	SERVER_ADDR = ":8080"
)

func main() {
	// Create required dependencies
	log := log.New(os.Stdout, "[products-service] ", log.LstdFlags)

	// Create handlers
	productsHandler := handler.NewProductsHandler(log)

	// Setup handlers with custom serve mux
	mux := http.NewServeMux()
	mux.Handle("/products", productsHandler)

	// Create custom server and setup configs
	server := &http.Server{
		Addr:         SERVER_ADDR, // Server host and port
		Handler:      mux,         // Custom mux
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

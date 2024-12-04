package main

import (
	"fmt"
	"go-store/controller"
	"go-store/repository"
	"go-store/router"
	"log"
)

func main() {
	// Initialize the repository
	repo := repository.NewProductRepository()

	// Initialize the controller
	productController := controller.NewProductController(repo)

	// Initialize the router
	r := router.NewRouter(productController)

	// Start the server
	addr := ":8080"
	fmt.Printf("Server is running on http://localhost%s\n", addr)
	if err := r.Run(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

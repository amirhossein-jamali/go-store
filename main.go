package main

import (
	"fmt"
	"go-store/controller"
	"go-store/repository"
	"go-store/router"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

func main() {
	// Initialize SQLite database
	db, err := gorm.Open(sqlite.Open("products.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect database: %v", err)
	}

	// Initialize repository
	repo := repository.NewGormProductRepository(db)

	// Initialize controller
	productController := controller.NewProductController(repo)

	// Initialize router
	r := router.NewRouter(productController)

	// Start the server
	addr := ":8080"
	fmt.Printf("Server is running on http://localhost%s\n", addr)
	if err := r.Run(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

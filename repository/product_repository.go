package repository

import (
	"errors"
	"go-store/model"
)

// InMemoryProductRepository implements a simple in-memory data storage for products.
type InMemoryProductRepository struct {
	products []model.Product
}

// NewInMemoryProductRepository initializes the repository with sample data.
func NewInMemoryProductRepository() *InMemoryProductRepository {
	return &InMemoryProductRepository{
		products: []model.Product{
			{ID: 1, Name: "Product A"},
			{ID: 2, Name: "Product B"},
		},
	}
}

// GetAllProducts retrieves all products.
func (repo *InMemoryProductRepository) GetAllProducts() []model.Product {
	return repo.products
}

// GetProductByID retrieves a product by its ID or returns an error if not found.
func (repo *InMemoryProductRepository) GetProductByID(id int) (*model.Product, error) {
	for _, product := range repo.products {
		if product.ID == id {
			return &product, nil
		}
	}
	return nil, errors.New("product not found")
}

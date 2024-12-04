package repository

import (
	"errors"
	"go-store/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// ProductRepository handles database operations for products.
type ProductRepository struct {
	db *gorm.DB
}

// NewProductRepository initializes the repository and sets up the database.
func NewProductRepository() *ProductRepository {
	// Connect to SQLite database
	db, err := gorm.Open(sqlite.Open("products.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Auto-migrate the Product model
	err = db.AutoMigrate(&model.Product{})
	if err != nil {
		panic("failed to migrate database")
	}

	return &ProductRepository{db: db}
}

// SaveProduct saves a product to the database.
func (repo *ProductRepository) SaveProduct(product *model.Product) error {
	return repo.db.Create(product).Error
}

// GetAllProducts retrieves all products from the database.
func (repo *ProductRepository) GetAllProducts() ([]model.Product, error) {
	var products []model.Product
	err := repo.db.Find(&products).Error
	return products, err
}

// GetProductByID retrieves a product by its ID.
func (repo *ProductRepository) GetProductByID(id uint) (*model.Product, error) {
	var product model.Product
	err := repo.db.First(&product, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &product, err
}

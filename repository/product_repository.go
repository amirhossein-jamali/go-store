package repository

import (
	"errors"
	"go-store/model"
	"gorm.io/gorm"
)

// ProductRepository defines the methods for product data management.
type ProductRepository interface {
	SaveProduct(product *model.Product) error
	GetAllProducts() ([]model.Product, error)
	GetProductByID(id uint) (*model.Product, error)
}

// GormProductRepository is an implementation of ProductRepository using GORM.
type GormProductRepository struct {
	db *gorm.DB
}

// NewGormProductRepository initializes the repository.
func NewGormProductRepository(db *gorm.DB) *GormProductRepository {
	return &GormProductRepository{db: db}
}

func (repo *GormProductRepository) SaveProduct(product *model.Product) error {
	return repo.db.Create(product).Error
}

func (repo *GormProductRepository) GetAllProducts() ([]model.Product, error) {
	var products []model.Product
	err := repo.db.Find(&products).Error
	return products, err
}

func (repo *GormProductRepository) GetProductByID(id uint) (*model.Product, error) {
	var product model.Product
	err := repo.db.First(&product, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, gorm.ErrRecordNotFound
	}
	return &product, err
}

package tests

import (
	"go-store/model"
	"go-store/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) (*gorm.DB, func()) {
	t.Helper()

	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to the in-memory database: %v", err)
	}

	if err := db.AutoMigrate(&model.Product{}); err != nil {
		t.Fatalf("Failed to migrate schema: %v", err)
	}

	cleanup := func() {
		dbSQL, err := db.DB()
		if err != nil {
			t.Errorf("Failed to get raw database instance: %v", err)
		}
		if err := dbSQL.Close(); err != nil {
			t.Errorf("Failed to close database: %v", err)
		}
	}

	return db, cleanup
}

func TestRepository_SaveProduct(t *testing.T) {
	t.Parallel() // Allow parallel execution

	db, cleanup := setupTestDB(t)
	defer cleanup()

	repo := repository.NewGormProductRepository(db)
	product := &model.Product{Name: "Test Product"}

	err := repo.SaveProduct(product)
	assert.NoError(t, err, "SaveProduct should not return an error")
	assert.NotZero(t, product.ID, "Product ID should not be zero after saving")
}

func TestRepository_GetAllProducts(t *testing.T) {
	t.Parallel()

	db, cleanup := setupTestDB(t)
	defer cleanup()

	repo := repository.NewGormProductRepository(db)

	// Add test data
	productsToAdd := []model.Product{
		{Name: "Product A"},
		{Name: "Product B"},
	}
	for _, p := range productsToAdd {
		if err := repo.SaveProduct(&p); err != nil {
			t.Fatalf("Failed to save product %v: %v", p, err)
		}
	}

	// Retrieve all products
	products, err := repo.GetAllProducts()
	assert.NoError(t, err, "GetAllProducts should not return an error")
	assert.Len(t, products, len(productsToAdd), "Number of retrieved products should match added products")

	names := []string{"Product A", "Product B"}
	for _, product := range products {
		assert.Contains(t, names, product.Name, "Retrieved product name should match added products")
	}
}

func TestRepository_GetProductByID(t *testing.T) {
	t.Parallel()

	db, cleanup := setupTestDB(t)
	defer cleanup()

	repo := repository.NewGormProductRepository(db)

	product := &model.Product{Name: "Product A"}
	if err := repo.SaveProduct(product); err != nil {
		t.Fatalf("Failed to save product: %v", err)
	}

	// Retrieve by ID
	result, err := repo.GetProductByID(product.ID)
	assert.NoError(t, err, "GetProductByID should not return an error")
	assert.NotNil(t, result, "Retrieved product should not be nil")
	assert.Equal(t, product.Name, result.Name, "Retrieved product name should match the saved product")
}

func TestRepository_GetProductByID_NotFound(t *testing.T) {
	t.Parallel()

	db, cleanup := setupTestDB(t)
	defer cleanup()

	repo := repository.NewGormProductRepository(db)

	result, err := repo.GetProductByID(999)
	assert.Error(t, err, "Expected an error for non-existent ID")
	assert.Nil(t, result, "Result should be nil for non-existent ID")
	assert.Equal(t, gorm.ErrRecordNotFound, err, "Error should be ErrRecordNotFound")
}

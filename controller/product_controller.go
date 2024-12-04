package controller

import (
	"go-store/model"
	"go-store/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ProductController handles product-related requests.
type ProductController struct {
	repo *repository.ProductRepository
}

// NewProductController creates a new instance of ProductController.
func NewProductController(repo *repository.ProductRepository) *ProductController {
	return &ProductController{repo: repo}
}

// CreateProduct handles the creation of a new product.
func (pc *ProductController) CreateProduct(c *gin.Context) {
	var input model.Product

	// Parse and bind JSON input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Validate the product
	if err := input.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Save the product to the database
	if err := pc.repo.SaveProduct(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product created successfully", "product": input})
}

// GetProducts retrieves all products.
func (pc *ProductController) GetProducts(c *gin.Context) {
	products, err := pc.repo.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve products"})
		return
	}

	c.JSON(http.StatusOK, products)
}

// GetProductByID retrieves a product by ID.
func (pc *ProductController) GetProductByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	product, err := pc.repo.GetProductByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve product"})
		return
	}

	if product == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}

	c.JSON(http.StatusOK, product)
}

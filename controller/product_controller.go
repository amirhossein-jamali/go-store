package controller

import (
	"go-store/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ProductController handles product-related API requests.
type ProductController struct {
	repo *repository.InMemoryProductRepository
}

// NewProductController creates a new instance of ProductController.
func NewProductController(repo *repository.InMemoryProductRepository) *ProductController {
	return &ProductController{repo: repo}
}

// GetProducts retrieves all products and sends them as JSON.
func (pc *ProductController) GetProducts(c *gin.Context) {
	products := pc.repo.GetAllProducts()
	c.JSON(http.StatusOK, products)
}

// GetProductByID retrieves a product by its ID and sends it as JSON.
func (pc *ProductController) GetProductByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	product, err := pc.repo.GetProductByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, product)
}

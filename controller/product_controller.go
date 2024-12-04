package controller

import (
	"github.com/gin-gonic/gin"
	"go-store/model"
	"go-store/repository"
	"net/http"
	"strconv"
)

// ProductController handles API requests for products.
type ProductController struct {
	repo repository.ProductRepository
}

// NewProductController initializes the controller.
func NewProductController(repo repository.ProductRepository) *ProductController {
	return &ProductController{repo: repo}
}

func (pc *ProductController) CreateProduct(c *gin.Context) {
	var input model.Product

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := input.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := pc.repo.SaveProduct(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product created successfully", "product": input})
}
func (pc *ProductController) GetProducts(c *gin.Context) {
	products, err := pc.repo.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve products"})
		return
	}

	c.JSON(http.StatusOK, products)
}

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

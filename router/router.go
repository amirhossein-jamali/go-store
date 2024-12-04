package router

import (
	"go-store/controller"

	"github.com/gin-gonic/gin"
)

// NewRouter creates a new router and sets up API routes.
func NewRouter(pc *controller.ProductController) *gin.Engine {
	r := gin.Default()

	// Product routes
	r.GET("/products", pc.GetProducts)
	r.GET("/products/:id", pc.GetProductByID)
	r.POST("/products", pc.CreateProduct)

	return r
}

package router

import (
	"go-store/controller"

	"github.com/gin-gonic/gin"
)

// NewRouter creates a new router and configures the API endpoints.
func NewRouter(pc *controller.ProductController) *gin.Engine {
	r := gin.Default()

	// Define the API endpoints
	r.GET("/products", pc.GetProducts)
	r.GET("/products/:id", pc.GetProductByID)

	return r
}

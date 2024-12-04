package tests

import (
	"encoding/json"
	"go-store/controller"
	"go-store/mocks"
	"go-store/model"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Helper function to unmarshal response body
func parseJSONResponse(body *httptest.ResponseRecorder, target interface{}) error {
	return json.Unmarshal(body.Body.Bytes(), target)
}

// Helper function to initialize router and controller
func setupTestRouter(mockRepo *mocks.ProductRepository) *gin.Engine {
	router := gin.Default()
	pc := controller.NewProductController(mockRepo)
	router.POST("/products", pc.CreateProduct)
	router.GET("/products", pc.GetProducts)
	router.GET("/products/:id", pc.GetProductByID)
	return router
}

func TestAPI_CreateProduct_Success(t *testing.T) {
	t.Parallel()

	mockRepo := mocks.NewProductRepository(t)
	mockRepo.On("SaveProduct", mock.Anything).Return(nil)

	router := setupTestRouter(mockRepo)

	reqBody := `{"name": "Test Product"}`
	req := httptest.NewRequest("POST", "/products", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code, "Response code should be 200 OK")

	var response map[string]interface{}
	err := parseJSONResponse(resp, &response)
	assert.NoError(t, err, "Response JSON should be parsable")
	assert.Equal(t, "Product created successfully", response["message"], "Response message mismatch")
	product := response["product"].(map[string]interface{})
	assert.Equal(t, "Test Product", product["name"], "Product name mismatch")
}

func TestAPI_CreateProduct_InvalidRequest(t *testing.T) {
	t.Parallel()

	mockRepo := mocks.NewProductRepository(t)
	router := setupTestRouter(mockRepo)

	reqBody := `{"invalid_field": "Invalid"}`
	req := httptest.NewRequest("POST", "/products", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code, "Response code should be 400 Bad Request")
}

func TestAPI_GetProducts_Success(t *testing.T) {
	t.Parallel()

	mockRepo := mocks.NewProductRepository(t)
	mockRepo.On("GetAllProducts").Return([]model.Product{
		{Name: "Product A"},
		{Name: "Product B"},
	}, nil)

	router := setupTestRouter(mockRepo)

	req := httptest.NewRequest("GET", "/products", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code, "Response code should be 200 OK")

	var response []model.Product
	err := parseJSONResponse(resp, &response)
	assert.NoError(t, err, "Response JSON should be parsable")
	assert.Len(t, response, 2, "Response should contain 2 products")
	assert.Equal(t, "Product A", response[0].Name, "First product name mismatch")
	assert.Equal(t, "Product B", response[1].Name, "Second product name mismatch")
}

func TestAPI_GetProducts_EmptyList(t *testing.T) {
	t.Parallel()

	mockRepo := mocks.NewProductRepository(t)
	mockRepo.On("GetAllProducts").Return([]model.Product{}, nil)

	router := setupTestRouter(mockRepo)

	req := httptest.NewRequest("GET", "/products", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code, "Response code should be 200 OK")

	var response []model.Product
	err := parseJSONResponse(resp, &response)
	assert.NoError(t, err, "Response JSON should be parsable")
	assert.Empty(t, response, "Response should contain no products")
}

func TestAPI_GetProductByID_Success(t *testing.T) {
	t.Parallel()

	mockRepo := mocks.NewProductRepository(t)
	mockRepo.On("GetProductByID", uint(1)).Return(&model.Product{Name: "Product A"}, nil)

	router := setupTestRouter(mockRepo)

	req := httptest.NewRequest("GET", "/products/1", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code, "Response code should be 200 OK")

	var product model.Product
	err := parseJSONResponse(resp, &product)
	assert.NoError(t, err, "Response JSON should be parsable")
	assert.Equal(t, "Product A", product.Name, "Product name mismatch")
}

func TestAPI_GetProductByID_NotFound(t *testing.T) {
	t.Parallel()

	mockRepo := mocks.NewProductRepository(t)
	mockRepo.On("GetProductByID", uint(999)).Return(nil, nil)

	router := setupTestRouter(mockRepo)

	req := httptest.NewRequest("GET", "/products/999", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusNotFound, resp.Code, "Response code should be 404 Not Found")
}

func TestAPI_GetProductByID_Error(t *testing.T) {
	t.Parallel()

	mockRepo := mocks.NewProductRepository(t)
	mockRepo.On("GetProductByID", uint(999)).Return(nil, mock.Anything)

	router := setupTestRouter(mockRepo)

	req := httptest.NewRequest("GET", "/products/999", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusInternalServerError, resp.Code, "Response code should be 500 Internal Server Error")
}

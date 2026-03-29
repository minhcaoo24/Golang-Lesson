package v1handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
}

func NewProductHandler() *ProductHandler {
	return &ProductHandler{}
}

// ===== Products =====

// GET /api/v1/Products
func (u *ProductHandler) GetProductsV1(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"respone": "Get all Products, List all product (v1)",
	})
}

// GET /api/v1/Products/:id
func (u *ProductHandler) GetProductsByIdV1(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"respone": "Okay, Get product by id (v1)",
	})
}

// POST /api/v1/Products
func (u *ProductHandler) PostProductsV1(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"respone": "Okay, Create product (v1)",
	})
}

// PUT /api/v1/Products/:id
func (u *ProductHandler) PutProductsV1(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"respone": "Okay, Update product (v1)",
	})
}

// DELETE /api/v1/Products/:id
func (u *ProductHandler) DeleteProductsV1(c *gin.Context) {
	c.JSON(http.StatusNoContent, gin.H{
		"respone": "Okay, delete product (v1)",
	})
}

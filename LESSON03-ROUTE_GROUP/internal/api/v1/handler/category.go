package v1handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var validCategory = map[string]bool{
	"C#":     true,
	"C++":    true,
	"Golang": true,
}

type CategoryHandler struct {
}

func NewCategoryHandler() *CategoryHandler {
	return &CategoryHandler{}
}

func (ca *CategoryHandler) GetCategoryByCategoryV1(c *gin.Context) {
	category := c.Param("category")

	if !validCategory[category] {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Category must be C#, C++ or Golang",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"respone":  "Category found",
		"category": category,
	})
}

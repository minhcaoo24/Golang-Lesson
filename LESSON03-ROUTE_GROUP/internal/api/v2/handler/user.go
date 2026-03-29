package v2handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

// ===== USERS =====

// GET /api/v1/users
func (u *UserHandler) GetUsersV1(c *gin.Context) {
	c.JSON(http.StatusOK, "Get all users, List all users (v1)")
}

// GET /api/v1/users/:id
func (u *UserHandler) GetUsersByIdV1(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"respone": "Okay, Get user by id (v2)",
	})
}

// POST /api/v1/users
func (u *UserHandler) PostUsersV1(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"respone": "Okay, Create user (v2)",
	})
}

// PUT /api/v1/users/:id
func (u *UserHandler) PutUsersV1(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"respone": "Okay, Update user (v2)",
	})
}

// DELETE /api/v1/users/:id
func (u *UserHandler) DeleteUsersV1(c *gin.Context) {
	c.JSON(http.StatusNoContent, gin.H{
		"respone": "Okay, delete user (v1)",
	})
}

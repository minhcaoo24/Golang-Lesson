package v1handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID must be a number",
		})
		return
	}

	if id < 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "ID cannot be negative",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"respone": "Okay, Get user by id (v1)",
		"id":      id,
	})
}

func (u *UserHandler) GetUsersByUuidV1(c *gin.Context) {
	uuidStr := c.Param("uuid")
	uuid, err := uuid.Parse(uuidStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid uuid, please re-check",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"respone": "Okay, Get user by uuid (v1)",
		"uuid":    uuid,
	})
}

// POST /api/v1/users
func (u *UserHandler) PostUsersV1(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"respone": "Okay, Create user (v1)",
	})
}

// PUT /api/v1/users/:id
func (u *UserHandler) PutUsersV1(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"respone": "Okay, Update user (v1)",
	})
}

// DELETE /api/v1/users/:id
func (u *UserHandler) DeleteUsersV1(c *gin.Context) {
	c.JSON(http.StatusNoContent, gin.H{
		"respone": "Okay, delete user (v1)",
	})
}

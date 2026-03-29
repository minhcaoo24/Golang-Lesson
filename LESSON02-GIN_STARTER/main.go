package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/demo", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Hello, Bitch",
		})
	})

	r.GET("/users", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"respone": "list of users",
		})
	})

	r.GET("/user/:user_id", func(ctx *gin.Context) {
		user_id := ctx.Param("user_id")
		ctx.JSON(http.StatusOK, gin.H{
			"data":    "Thong tin user",
			"user id": user_id,
		})
	})

	r.GET("/products", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"respone": "List of products",
		})
	})

	r.GET("/product/:product_name", func(ctx *gin.Context) {
		product_name := ctx.Param("product_name")

		price := ctx.Query("price")
		color := ctx.Query("color")

		ctx.JSON(http.StatusOK, gin.H{
			"data":         "Thong tin product",
			"product name": product_name,
			"price":        price,
			"color":        color,
		})
	})
	r.Run(":8080")
}

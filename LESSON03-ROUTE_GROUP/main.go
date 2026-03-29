package main

import (
	v1handler "route_group/internal/api/v1/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		user := v1.Group("users")
		{
			userHandlerV1 := v1handler.NewUserHandler()
			user.GET("", userHandlerV1.GetUsersV1)
			user.GET("/:id", userHandlerV1.GetUsersByIdV1)
			user.POST("", userHandlerV1.PostUsersV1)
			user.PUT(":id", userHandlerV1.PutUsersV1)
			user.DELETE("/:id", userHandlerV1.DeleteUsersV1)
		}

		product := v1.Group("products")
		{
			productHandlerV1 := v1handler.NewProductHandler()
			product.GET("", productHandlerV1.GetProductsV1)
			product.GET("/:id", productHandlerV1.GetProductsByIdV1)
			product.POST("", productHandlerV1.PostProductsV1)
			product.PUT("/:id", productHandlerV1.PutProductsV1)
			product.DELETE("/:id", productHandlerV1.DeleteProductsV1)
		}
	}

	r.Run(":8080")
}

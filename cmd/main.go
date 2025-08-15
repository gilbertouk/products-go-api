package main

import (
	"go-api/controller"
	"go-api/db"
	"go-api/repository"
	"go-api/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the Gin router
	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	// Initialize the product repository
	ProductRepository := repository.NewProductRepository(dbConnection)

	// Initialize the product use case
	ProductUseCase := usecase.NewProductUseCase(ProductRepository)

	// Initialize the product controller
	ProductController := controller.NewProductController(ProductUseCase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Define the product routes
	server.GET("/products", ProductController.GetProducts)
	server.POST("/products", ProductController.CreateProduct)
	server.GET("/products/:productId", ProductController.GetProductByID)
	server.PUT("/products/:productId", ProductController.UpdateProduct)

	server.Run(":8000")
}

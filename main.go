package main

import (
	"example/Crud/config"
	"example/Crud/controllers"
	"example/Crud/logger"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//@title CRUD API
// @version 1.0
// @description This is a CRUD API documentation
// @termsOfService http://swagger.io/terms/

//@host localhost:8080
//@BasePath	/

func main() {
	logger.CustomLogger.Info("Starting Application")

	config.ConnectDatabase()

	router := gin.Default() //default middleware
	//swagger documentation route
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.POST("/items", controllers.CreateItem)
	router.GET("/items", controllers.GetItems)
	router.GET("/items/:id", controllers.GetItemByID)
	router.PUT("/items/:id", controllers.UpdateItem)
	router.DELETE("/items/:id", controllers.DeleteItem)

	logger.CustomLogger.Info("Application running on Port 8080")
	router.Run(":8080")
}

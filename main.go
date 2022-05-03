package main

import (
	"github.com/gin-gonic/gin"

	"restapi/controllers"
	"restapi/models"
)

func main() {
	router := gin.Default()

	models.ConnectDatabase()

	router.GET("/books", controllers.FindBooks)
	router.GET("/books/:id", controllers.FindBook)
	router.POST("/books", controllers.CreateBook)
	router.PATCH("/books/:id", controllers.UpdateBook)
	router.DELETE("/books/:id", controllers.DeleteBook)

	router.Run()
}

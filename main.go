package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go-crud.com/controllers"
	"go-crud.com/models"
)

func main() {
	models.ConnectDatabase()
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})
	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)
	r.GET("/books/:id", controllers.FindBook)
	r.POST("/books/delete/:id", controllers.DeleteBook)
	r.Run()
}

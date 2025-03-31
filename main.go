package main

import (
	"book-management-api/configs"
	"book-management-api/controllers"
	db "book-management-api/internals/dbservice"
	"book-management-api/internals/services"
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	err := configs.NewApplicationConfig(context.Background())
	if err != nil {
		fmt.Println("Error in Appconfig:", err)
	}

	dbService := db.NewDbService(configs.AppConfig.DbClient)

	r := gin.Default()

	bookService := services.NewBookService(dbService)
	bookController := controllers.NewBookController(bookService)

	r.GET("/books", bookController.GetBooks)
	r.GET("/books/:id", bookController.GetBookById)
	r.POST("/books", bookController.CreateBook)
	r.PUT("/books/:id", bookController.UpdateBook)
	r.DELETE("/books/:id", bookController.DeleteBook)

	r.Run(":" + configs.AppConfig.HttpPort)
}

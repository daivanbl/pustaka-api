package main

import (
	"pustaka-api/config"
	"pustaka-api/route"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.ConnectDatabase()
	// bookRepository := domain.NewRepository(db)
	// bookService := service.NewService(bookRepository)
	// bookHandler := handler.NewBookHandler(bookService)

	router := gin.Default()

	route.ApiRoute(router, db)

	// books := router.Group("/books")

	// books.GET("/", bookHandler.BooksGetAll)
	// books.GET("/:id", bookHandler.BooksGetById)
	// books.GET("/query", handler.QueryHandler)
	// books.POST("/", bookHandler.InsertBook)
	// books.POST("/:id", bookHandler.UpdateBook)

	router.Run(":8888")
}

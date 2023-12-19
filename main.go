package main

import (
	"fmt"
	"log"
	"pustaka-api/domain"
	"pustaka-api/handler"
	"pustaka-api/service"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/pustaka-api?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB connection error")
	} else {
		fmt.Println("DB connected")
	}

	db.AutoMigrate(&domain.Book{})

	bookRepository := domain.NewRepository(db)
	bookService := service.NewService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)

	router := gin.Default()

	books := router.Group("/books")

	books.GET("/", bookHandler.BooksGetAll)
	books.GET("/:id", bookHandler.BooksGetById)
	books.GET("/query", handler.QueryHandler)
	books.POST("/", bookHandler.InsertBook)
	books.POST("/:id", bookHandler.UpdateBook)

	router.Run(":8888")
}

package route

import (
	"pustaka-api/domain"
	"pustaka-api/handler"
	"pustaka-api/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func BookRoute(books *gin.RouterGroup, db *gorm.DB) {

	bookRepository := domain.NewRepository(db)
	bookService := service.NewService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)

	books.GET("/", bookHandler.BooksGetAll)
	books.GET("/:id", bookHandler.BooksGetById)
	books.GET("/query", handler.QueryHandler)
	books.POST("/", bookHandler.InsertBook)
	books.POST("/:id", bookHandler.UpdateBook)

}

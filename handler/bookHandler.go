package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"

	"pustaka-api/domain"
	"pustaka-api/service"
)

type bookHandler struct {
	bookService service.BookService
}

func NewBookHandler(bookService service.BookService) *bookHandler {
	return &bookHandler{bookService}
}

func (handler *bookHandler) BooksGetAll(c *gin.Context) {
	var response []domain.BookResponse
	books, _ := handler.bookService.FindAll()
	for _, book := range books {
		response = append(response, constructBookResponse(book))
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"data":   response,
	})
}

func (handler *bookHandler) BooksGetById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	book, _ := handler.bookService.FindById(id)

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"data":   constructBookResponse(book),
	})
}

func QueryHandler(c *gin.Context) {
	title := c.Query("title")

	c.JSON(http.StatusOK, gin.H{"title": title})
}

func (handler *bookHandler) InsertBook(c *gin.Context) {
	var bookRequest domain.BookRequest

	err := c.ShouldBindJSON(&bookRequest)
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"errors": errorMessages,
		})
		return
	}

	newBook, err := handler.bookService.Create(bookRequest)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"result":  newBook,
		"message": "Success insert new row",
	})
}

func (handler *bookHandler) UpdateBook(c *gin.Context) {
	var bookRequest domain.BookRequest
	id, _ := strconv.Atoi(c.Param("id"))

	err := c.ShouldBindJSON(&bookRequest)
	if err != nil {
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on field %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"errors": errorMessages,
		})
		return
	}

	updatedBook, err := handler.bookService.Update(bookRequest, id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": 400,
			"errors": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"result":  constructBookResponse(updatedBook),
		"message": "Success Update Book with ID :" + c.Param("id"),
	})
}

func constructBookResponse(book domain.Book) domain.BookResponse {
	return domain.BookResponse{
		ID:          book.ID,
		Title:       book.Title,
		Description: book.Description,
		AuthorName:  book.AuthorName,
		Price:       book.Price,
		Rating:      book.Rating,
	}
}

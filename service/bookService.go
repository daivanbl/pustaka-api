package service

import (
	"pustaka-api/domain"
	"time"
)

type BookService interface {
	FindAll() ([]domain.Book, error)
	FindById(ID int) (domain.Book, error)
	Create(book domain.BookRequest) (domain.Book, error)
	Update(book domain.BookRequest, ID int) (domain.Book, error)
}

type service struct {
	repository domain.BookRepository
}

func NewService(repository domain.BookRepository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]domain.Book, error) {
	books, err := s.repository.FindAll()
	return books, err
}

func (s *service) FindById(ID int) (domain.Book, error) {
	book, err := s.repository.FindById(ID)
	return book, err
}

func (s *service) Create(bookRequest domain.BookRequest) (domain.Book, error) {
	price, _ := bookRequest.Price.Int64()
	book := domain.Book{
		Title:       bookRequest.Title,
		Description: bookRequest.Description,
		AuthorName:  bookRequest.AuthorName,
		Price:       int(price),
		Rating:      bookRequest.Rating,
	}

	newBook, err := s.repository.Create(book)

	return newBook, err
}

func (s *service) Update(bookRequest domain.BookRequest, ID int) (domain.Book, error) {
	book, err := s.repository.FindById(ID)
	price, _ := bookRequest.Price.Int64()

	book.Title = bookRequest.Title
	book.Description = bookRequest.Description
	book.AuthorName = bookRequest.AuthorName
	book.Price = int(price)
	book.Rating = bookRequest.Rating
	book.UpdatedAt = time.Now()

	updatedBook, err := s.repository.Update(book, ID)

	return updatedBook, err

}

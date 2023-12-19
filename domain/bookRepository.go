package domain

import "gorm.io/gorm"

type BookRepository interface {
	FindAll() ([]Book, error)
	FindById(ID int) (Book, error)
	Create(book Book) (Book, error)
	Update(book Book, ID int) (Book, error)
}

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) FindAll() ([]Book, error) {
	var books []Book
	err := r.db.Find(&books).Error
	return books, err
}

func (r *Repository) FindById(ID int) (Book, error) {
	var book Book
	err := r.db.Find(&book, ID).Error
	return book, err
}

func (r *Repository) Create(book Book) (Book, error) {
	err := r.db.Create(&book).Error
	return book, err
}

func (r *Repository) Update(book Book, ID int) (Book, error) {
	err := r.db.Save(&book).Where("id = ?", ID).Error
	return book, err
}

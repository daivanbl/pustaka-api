package domain

import "time"

type Book struct {
	ID          int
	Title       string
	Description string
	AuthorName  string
	Price       int
	Rating      int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

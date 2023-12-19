package domain

import "encoding/json"

type BookRequest struct {
	Title       string      `json:"title" binding:"required"`
	Price       json.Number `json:"price" binding:"required,number"`
	Description string      `json:"description" binding:"required"`
	AuthorName  string      `json:"authorName" binding:"required"`
	Rating      int         `json:"rating" binding:"required,number"`
}

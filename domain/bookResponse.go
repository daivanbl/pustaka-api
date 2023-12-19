package domain

type BookResponse struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	AuthorName  string `json:"authorName"`
	Price       int    `json:"price"`
	Rating      int    `json:"rating"`
}

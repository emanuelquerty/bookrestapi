package models

type Book struct {
	BookID        int    `json:"book_id,omitempty"`
	Title         string `json:"title,omitempty"`
	PublishedYear int    `json:"published_year,omitempty"`
	AuthorID      int    `json:"author_id,omitempty"`
}

type BookResponse struct {
	Status     string `json:"status,omitempty"`
	StatusCode int    `json:"status_code,omitempty"`
	Data       []Book `json:"data,omitempty"`
}

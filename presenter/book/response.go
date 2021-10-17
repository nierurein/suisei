package book

import (
	"time"

	"github.com/daniel5u/suisei/domain/book"
)

type Response struct {
	ID              int       `json:"id"`
	PublisherID     int       `json:"publisher_id"`
	CategoryID      int       `json:"category_id"`
	ISBN            string    `json:"isbn"`
	Title           string    `json:"title"`
	Price           int       `json:"price"`
	PublicationYear int       `json:"publication_year"`
	PageCount       int       `json:"page_count"`
	Description     string    `json:"description"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

func domainToResponse(bookDomain book.Domain) Response {
	return Response{
		ID:              bookDomain.ID,
		PublisherID:     bookDomain.PublisherID,
		CategoryID:      bookDomain.CategoryID,
		ISBN:            bookDomain.ISBN,
		Title:           bookDomain.Title,
		Price:           bookDomain.Price,
		PublicationYear: bookDomain.PublicationYear,
		PageCount:       bookDomain.PageCount,
		Description:     bookDomain.Description,
		CreatedAt:       bookDomain.CreatedAt,
		UpdatedAt:       bookDomain.UpdatedAt,
	}
}

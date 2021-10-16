package book

import "github.com/daniel5u/suisei/domain/book"

type Request struct {
	PublisherID     int    `json:"publisher_id"`
	CategoryID      int    `json:"category_id"`
	ISBN            string `json:"isbn"`
	Title           string `json:"title"`
	Price           int    `json:"price"`
	PublicationYear int    `json:"publication_year"`
	PageCount       int    `json:"page_count"`
	Description     string `json:"description"`
}

func requestToDomain(request Request) book.Domain {
	return book.Domain{
		PublisherID:     request.PublisherID,
		CategoryID:      request.CategoryID,
		ISBN:            request.ISBN,
		Title:           request.Title,
		Price:           request.Price,
		PublicationYear: request.PublicationYear,
		PageCount:       request.PageCount,
		Description:     request.Description,
	}
}

package openlibrary

import (
	"math/rand"

	"github.com/daniel5u/suisei/domain/openlibrary"
)

type Response struct {
	Publishers      []string            `json:"publishers"`
	Category        []string            `json:"subjects"`
	ISBN            []string            `json:"isbn_10"`
	Authors         []map[string]string `json:"authors"`
	Title           string              `json:"title"`
	Price           int                 `json:"price"`
	PublicationDate string              `json:"publish_date"`
	PageCount       int                 `json:"number_of_pages"`
	Description     string              `json:"description"`
}

func responseToDomain(response Response, authors []string, publicationYear int) openlibrary.Domain {
	return openlibrary.Domain{
		Publisher:       response.Publishers[0],
		Category:        response.Category[0],
		ISBN:            response.ISBN[0],
		Authors:         authors,
		Title:           response.Title,
		Price:           1 + rand.Intn(99),
		PublicationYear: publicationYear,
		PageCount:       response.PageCount,
		Description:     response.Description,
	}
}

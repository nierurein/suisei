package book

import (
	"github.com/daniel5u/suisei/domain/book"
	"github.com/daniel5u/suisei/repository/postgresql/bookauthor"
	"github.com/daniel5u/suisei/repository/postgresql/booktransaction"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	PublisherID      int
	CategoryID       int
	ISBN             string
	Title            string
	Price            int
	PublicationYear  int
	PageCount        int
	Description      string
	BookAuthors      []bookauthor.BookAuthor
	BookTransactions []booktransaction.BookTransaction
}

func repositoryToDomain(bookRepository Book) book.Domain {
	return book.Domain{
		ID:              int(bookRepository.ID),
		PublisherID:     bookRepository.PublisherID,
		CategoryID:      bookRepository.CategoryID,
		ISBN:            bookRepository.ISBN,
		Title:           bookRepository.Title,
		Price:           bookRepository.Price,
		PublicationYear: bookRepository.PublicationYear,
		PageCount:       bookRepository.PageCount,
		Description:     bookRepository.Description,
		CreatedAt:       bookRepository.CreatedAt,
		UpdatedAt:       bookRepository.UpdatedAt,
	}
}

func domainToRepository(bookDomain book.Domain) Book {
	return Book{
		PublisherID:     bookDomain.PublisherID,
		CategoryID:      bookDomain.CategoryID,
		ISBN:            bookDomain.ISBN,
		Title:           bookDomain.Title,
		Price:           bookDomain.Price,
		PublicationYear: bookDomain.PublicationYear,
		PageCount:       bookDomain.PageCount,
		Description:     bookDomain.Description,
	}
}

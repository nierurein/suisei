package bookauthor

import (
	"github.com/daniel5u/suisei/domain/bookauthor"
	"gorm.io/gorm"
)

type BookAuthor struct {
	gorm.Model
	BookID   int `gorm:"primaryKey"`
	AuthorID int `gorm:"primaryKey"`
}

func repositoryToDomain(bookauthorRepository BookAuthor) bookauthor.Domain {
	return bookauthor.Domain{
		ID:        int(bookauthorRepository.ID),
		BookID:    bookauthorRepository.BookID,
		AuthorID:  bookauthorRepository.AuthorID,
		CreatedAt: bookauthorRepository.CreatedAt,
		UpdatedAt: bookauthorRepository.UpdatedAt,
	}
}

func domainToRepository(bookauthorDomain bookauthor.Domain) BookAuthor {
	return BookAuthor{
		BookID:   bookauthorDomain.BookID,
		AuthorID: bookauthorDomain.AuthorID,
	}
}

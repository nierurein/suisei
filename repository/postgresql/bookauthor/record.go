package bookauthor

import (
	"github.com/daniel5u/suisei/domain/bookauthor"
)

type BookAuthor struct {
	BookID   int `gorm:"primaryKey"`
	AuthorID int `gorm:"primaryKey"`
}

func repositoryToDomain(bookauthorRepository BookAuthor) bookauthor.Domain {
	return bookauthor.Domain{
		BookID:   bookauthorRepository.BookID,
		AuthorID: bookauthorRepository.AuthorID,
	}
}

func domainToRepository(bookauthorDomain bookauthor.Domain) BookAuthor {
	return BookAuthor{
		BookID:   bookauthorDomain.BookID,
		AuthorID: bookauthorDomain.AuthorID,
	}
}

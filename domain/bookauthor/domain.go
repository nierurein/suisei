package bookauthor

import (
	"time"
)

type Domain struct {
	ID        int
	BookID    int
	AuthorID  int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Service interface {
	Store(bookauthorDomain Domain) (Domain, error)
	DeleteByBookID(bookid int) error
	DeleteByAuthorID(authorid int) error
}

type Repository interface {
	Store(bookauthorDomain Domain) (Domain, error)
	DeleteByBookID(bookid int) error
	DeleteByAuthorID(authorid int) error
}

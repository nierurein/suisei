package book

import (
	"time"
)

type Domain struct {
	ID              int
	PublisherID     int
	CategoryID      int
	ISBN            string
	Title           string
	Price           int
	PublicationYear int
	PageCount       int
	Description     string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type Service interface {
	Fetch() ([]Domain, error)
	GetByID(id int) (Domain, error)
	GetByTitle(title string) (Domain, error)
	Update(bookDomain Domain, id int) (Domain, error)
	Store(bookDomain Domain) (Domain, error)
	Delete(id int) error
}

type Repository interface {
	Fetch() ([]Domain, error)
	GetByID(id int) (Domain, error)
	GetByTitle(title string) (Domain, error)
	Update(bookDomain Domain, id int) (Domain, error)
	Store(bookDomain Domain) (Domain, error)
	Delete(id int) error
}

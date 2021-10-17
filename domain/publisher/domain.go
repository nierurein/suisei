package publisher

import (
	"time"
)

type Domain struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Service interface {
	Fetch() ([]Domain, error)
	GetByID(id int) (Domain, error)
	GetByName(name string) (Domain, error)
	Store(publisherDomain Domain) (Domain, error)
}

type Repository interface {
	Fetch() ([]Domain, error)
	GetByID(id int) (Domain, error)
	GetByName(name string) (Domain, error)
	Store(publisherDomain Domain) (Domain, error)
}

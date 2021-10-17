package user

import (
	"time"
)

type Domain struct {
	ID        int
	Name      string
	Email     string
	Password  string
	Balance   int
	Token     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Service interface {
	Fetch() ([]Domain, error)
	GetByID(id int) (Domain, error)
	Update(userDomain Domain, id int) (Domain, error)
	UpdateBalance(userDomain Domain, id int) error
	Store(userDomain Domain) (Domain, error)
	Delete(id int) error
}

type Repository interface {
	Fetch() ([]Domain, error)
	GetByID(id int) (Domain, error)
	Update(userDomain Domain, id int) (Domain, error)
	UpdateBalance(userDomain Domain, id int) error
	Store(userDomain Domain) (Domain, error)
	Delete(id int) error
}

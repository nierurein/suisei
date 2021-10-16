package transaction

import (
	"time"
)

type Domain struct {
	ID         int
	UserID     int
	TotalPrice int
	Status     int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Service interface {
	Fetch() ([]Domain, error)
	GetByID(id int) (Domain, error)
	Update(transactionDomain Domain, id int) (Domain, error)
	Store(transactionDomain Domain) (Domain, error)
}

type Repository interface {
	Fetch() ([]Domain, error)
	GetByID(id int) (Domain, error)
	Update(transactionDomain Domain, id int) (Domain, error)
	Store(transactionDomain Domain) (Domain, error)
}

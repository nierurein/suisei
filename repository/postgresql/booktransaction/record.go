package booktransaction

import (
	"github.com/daniel5u/suisei/domain/booktransaction"
)

type BookTransaction struct {
	TransactionID int `gorm:"primaryKey"`
	BookID        int `gorm:"primaryKey"`
	Quantity      int
	PricePerUnit  int
}

func repositoryToDomain(booktransactionRepository BookTransaction) booktransaction.Domain {
	return booktransaction.Domain{
		TransactionID: booktransactionRepository.TransactionID,
		BookID:        booktransactionRepository.BookID,
		Quantity:      booktransactionRepository.Quantity,
		PricePerUnit:  booktransactionRepository.PricePerUnit,
	}
}

func domainToRepository(booktransactionDomain booktransaction.Domain) BookTransaction {
	return BookTransaction{
		TransactionID: booktransactionDomain.TransactionID,
		BookID:        booktransactionDomain.BookID,
		Quantity:      booktransactionDomain.Quantity,
		PricePerUnit:  booktransactionDomain.PricePerUnit,
	}
}

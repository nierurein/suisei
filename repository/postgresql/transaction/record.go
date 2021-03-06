package transaction

import (
	"github.com/daniel5u/suisei/domain/transaction"
	"github.com/daniel5u/suisei/repository/postgresql/booktransaction"
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	UserID           int
	TotalPrice       int
	Status           int
	BookTransactions []booktransaction.BookTransaction
}

func repositoryToDomain(transactionRepository Transaction) transaction.Domain {
	return transaction.Domain{
		ID:         int(transactionRepository.ID),
		UserID:     transactionRepository.UserID,
		TotalPrice: transactionRepository.TotalPrice,
		Status:     transactionRepository.Status,
		CreatedAt:  transactionRepository.CreatedAt,
		UpdatedAt:  transactionRepository.UpdatedAt,
	}
}

func domainToRepository(transactionDomain transaction.Domain) Transaction {
	return Transaction{
		UserID:     transactionDomain.UserID,
		TotalPrice: transactionDomain.TotalPrice,
		Status:     transactionDomain.Status,
	}
}

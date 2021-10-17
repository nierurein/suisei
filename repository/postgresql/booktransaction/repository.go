package booktransaction

import (
	"github.com/daniel5u/suisei/domain/booktransaction"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) booktransaction.Repository {
	return &Repository{
		DB: db,
	}
}

func (booktransactionRepository *Repository) Store(booktransactionDomain booktransaction.Domain) (booktransaction.Domain, error) {
	var booktransactionRecord BookTransaction = domainToRepository(booktransactionDomain)

	err := booktransactionRepository.DB.Create(&booktransactionRecord).Error
	if err != nil {
		return booktransaction.Domain{}, err
	}

	booktransactionDomainAfter := repositoryToDomain(booktransactionRecord)

	return booktransactionDomainAfter, nil
}

func (booktransactionRepository *Repository) DeleteByTransactionID(transactionid int) error {
	var booktransactionRecord BookTransaction

	// permanent delete
	err := booktransactionRepository.DB.Unscoped().Where("transaction_id = ?", transactionid).Delete(&booktransactionRecord).Error
	if err != nil {
		return err
	}

	return nil
}

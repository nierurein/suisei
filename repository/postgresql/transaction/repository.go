package transaction

import (
	"github.com/daniel5u/suisei/domain/transaction"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) transaction.Repository {
	return &Repository{
		DB: db,
	}
}

func (transactionRepository *Repository) Fetch() ([]transaction.Domain, error) {
	var transactionRecords []Transaction
	var transactionDomains []transaction.Domain

	err := transactionRepository.DB.Find(&transactionRecords).Error
	if err != nil {
		return []transaction.Domain{}, err
	}

	for _, transactionRecord := range transactionRecords {
		transactionDomains = append(transactionDomains, repositoryToDomain(transactionRecord))
	}

	return transactionDomains, nil
}

func (transactionRepository *Repository) GetByID(id int) (transaction.Domain, error) {
	var transactionRecord Transaction

	err := transactionRepository.DB.Where("id = ?", id).First(&transactionRecord).Error
	if err != nil {
		return transaction.Domain{}, err
	}

	transactionDomain := repositoryToDomain(transactionRecord)

	return transactionDomain, nil
}

func (transactionRepository *Repository) Update(transactionDomain transaction.Domain, id int) (transaction.Domain, error) {
	var transactionRecord Transaction = domainToRepository(transactionDomain)
	var transactionRecordAfter Transaction

	err := transactionRepository.DB.Where("id = ?", id).Updates(&transactionRecord).Error
	if err != nil {
		return transaction.Domain{}, err
	}

	// get updated row
	err = transactionRepository.DB.Where("id = ?", id).First(&transactionRecordAfter).Error
	if err != nil {
		return transaction.Domain{}, err
	}

	transactionDomainAfter := repositoryToDomain(transactionRecordAfter)

	return transactionDomainAfter, nil
}

func (transactionRepository *Repository) Store(transactionDomain transaction.Domain) (transaction.Domain, error) {
	var transactionRecord Transaction = domainToRepository(transactionDomain)

	err := transactionRepository.DB.Create(&transactionRecord).Error
	if err != nil {
		return transaction.Domain{}, err
	}

	transactionDomainAfter := repositoryToDomain(transactionRecord)

	return transactionDomainAfter, nil
}

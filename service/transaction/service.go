package transaction

import (
	"github.com/daniel5u/suisei/domain/transaction"
)

type Service struct {
	repository transaction.Repository
}

func NewService(transactionRepository transaction.Repository) transaction.Service {
	return &Service{
		repository: transactionRepository,
	}
}

func (transactionService *Service) Fetch() ([]transaction.Domain, error) {
	result, err := transactionService.repository.Fetch()
	if err != nil {
		return []transaction.Domain{}, err
	}

	return result, nil
}

func (transactionService *Service) GetByID(id int) (transaction.Domain, error) {
	result, err := transactionService.repository.GetByID(id)
	if err != nil {
		return transaction.Domain{}, err
	}

	return result, nil
}

func (transactionService *Service) Update(transactionDomain transaction.Domain, id int) (transaction.Domain, error) {
	result, err := transactionService.repository.Update(transactionDomain, id)
	if err != nil {
		return transaction.Domain{}, err
	}

	return result, nil
}

func (transactionService *Service) Store(transactionDomain transaction.Domain) (transaction.Domain, error) {
	result, err := transactionService.repository.Store(transactionDomain)
	if err != nil {
		return transaction.Domain{}, err
	}

	return result, nil
}

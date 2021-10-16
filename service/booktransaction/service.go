package booktransaction

import (
	"github.com/daniel5u/suisei/domain/booktransaction"
)

type Service struct {
	repository booktransaction.Repository
}

func NewService(booktransactionRepository booktransaction.Repository) booktransaction.Service {
	return &Service{
		repository: booktransactionRepository,
	}
}

func (booktransactionService *Service) Store(booktransactionDomain booktransaction.Domain) (booktransaction.Domain, error) {
	result, err := booktransactionService.repository.Store(booktransactionDomain)
	if err != nil {
		return booktransaction.Domain{}, err
	}

	return result, nil
}

func (booktransactionService *Service) DeleteByTransactionID(transactionid int) error {
	err := booktransactionService.repository.DeleteByTransactionID(transactionid)
	if err != nil {
		return err
	}

	return nil
}

func (booktransactionService *Service) DeleteByBookID(bookid int) error {
	err := booktransactionService.repository.DeleteByBookID(bookid)
	if err != nil {
		return err
	}

	return nil
}

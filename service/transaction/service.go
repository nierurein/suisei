package transaction

import (
	"errors"

	"github.com/daniel5u/suisei/app/constant"
	"github.com/daniel5u/suisei/domain/transaction"
	"github.com/daniel5u/suisei/domain/user"
)

type Service struct {
	repository  transaction.Repository
	userService user.Service
}

func NewService(transactionRepository transaction.Repository, us user.Service) transaction.Service {
	return &Service{
		repository:  transactionRepository,
		userService: us,
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
	var transactionDomainBefore transaction.Domain
	var result transaction.Domain
	var userDomain user.Domain
	var err error

	transactionDomainBefore, err = transactionService.repository.GetByID(id)
	if err != nil {
		return transaction.Domain{}, err
	}

	// transaction status from PENDING to ACCEPTED
	if transactionDomainBefore.Status == constant.STATUS_PENDING && transactionDomain.Status == constant.STATUS_ACCEPTED {
		userDomain, err = transactionService.userService.GetByID(transactionDomainBefore.UserID)
		if err != nil {
			return transaction.Domain{}, err
		}

		if userDomain.Balance < transactionDomainBefore.TotalPrice {
			return transaction.Domain{}, errors.New("insufficient funds")
		}

		userDomain.Balance -= transactionDomainBefore.TotalPrice

		err = transactionService.userService.UpdateBalance(userDomain, userDomain.ID)
		if err != nil {
			return transaction.Domain{}, err
		}
	}

	result, err = transactionService.repository.Update(transactionDomain, id)
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

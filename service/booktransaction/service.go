package booktransaction

import (
	"errors"

	"github.com/daniel5u/suisei/app/constant"
	"github.com/daniel5u/suisei/domain/book"
	"github.com/daniel5u/suisei/domain/booktransaction"
	"github.com/daniel5u/suisei/domain/transaction"
)

type Service struct {
	repository         booktransaction.Repository
	transactionService transaction.Service
	bookService        book.Service
}

func NewService(booktransactionRepository booktransaction.Repository, ts transaction.Service, bs book.Service) booktransaction.Service {
	return &Service{
		repository:         booktransactionRepository,
		transactionService: ts,
		bookService:        bs,
	}
}

func (booktransactionService *Service) StoreBatch(booktransactionDomains []booktransaction.Domain, transactionid int) error {
	var bookDomain book.Domain
	var err error
	var totalPrice int

	transactionDomain, err := booktransactionService.transactionService.GetByID(transactionid)
	if err != nil {
		return err
	}

	// check transaction status
	if transactionDomain.Status != constant.STATUS_PENDING {
		return errors.New("transaction is locked")
	}

	// delete old booktransaction
	err = booktransactionService.repository.DeleteByTransactionID(transactionid)
	if err != nil {
		return err
	}

	for _, booktransactionDomain := range booktransactionDomains {
		// update TransactionID
		booktransactionDomain.TransactionID = transactionid

		// update PricePerUnit
		bookDomain, err = booktransactionService.bookService.GetByID(booktransactionDomain.BookID)
		if err != nil {
			return err
		}

		booktransactionDomain.PricePerUnit = bookDomain.Price

		_, err = booktransactionService.repository.Store(booktransactionDomain)
		if err != nil {
			return err
		}

		totalPrice += booktransactionDomain.PricePerUnit * booktransactionDomain.Quantity
	}

	// update transaction totalPrice
	transactionDomain.TotalPrice = totalPrice
	_, err = booktransactionService.transactionService.Update(transactionDomain, transactionid)
	if err != nil {
		return err
	}

	return nil
}

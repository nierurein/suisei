package book

import (
	"github.com/daniel5u/suisei/domain/book"
)

type Service struct {
	repository book.Repository
}

func NewService(bookRepository book.Repository) book.Service {
	return &Service{
		repository: bookRepository,
	}
}

func (bookService *Service) Fetch() ([]book.Domain, error) {
	result, err := bookService.repository.Fetch()
	if err != nil {
		return []book.Domain{}, err
	}

	return result, nil
}

func (bookService *Service) GetByID(id int) (book.Domain, error) {
	result, err := bookService.repository.GetByID(id)
	if err != nil {
		return book.Domain{}, err
	}

	return result, nil
}

func (bookService *Service) GetByTitle(title string) (book.Domain, error) {
	result, err := bookService.repository.GetByTitle(title)
	if err != nil {
		return book.Domain{}, err
	}

	return result, nil
}

func (bookService *Service) Update(bookDomain book.Domain, id int) (book.Domain, error) {
	result, err := bookService.repository.Update(bookDomain, id)
	if err != nil {
		return book.Domain{}, err
	}

	return result, nil
}

func (bookService *Service) Store(bookDomain book.Domain) (book.Domain, error) {
	result, err := bookService.repository.Store(bookDomain)
	if err != nil {
		return book.Domain{}, err
	}

	return result, nil
}

func (bookService *Service) Delete(id int) error {
	err := bookService.repository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

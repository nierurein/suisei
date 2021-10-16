package bookauthor

import (
	"github.com/daniel5u/suisei/domain/bookauthor"
)

type Service struct {
	repository bookauthor.Repository
}

func NewService(bookauthorRepository bookauthor.Repository) bookauthor.Service {
	return &Service{
		repository: bookauthorRepository,
	}
}

func (bookauthorService *Service) Store(bookauthorDomain bookauthor.Domain) (bookauthor.Domain, error) {
	result, err := bookauthorService.repository.Store(bookauthorDomain)
	if err != nil {
		return bookauthor.Domain{}, err
	}

	return result, nil
}

func (bookauthorService *Service) DeleteByBookID(bookid int) error {
	err := bookauthorService.repository.DeleteByBookID(bookid)
	if err != nil {
		return err
	}

	return nil
}

func (bookauthorService *Service) DeleteByAuthorID(authorid int) error {
	err := bookauthorService.repository.DeleteByAuthorID(authorid)
	if err != nil {
		return err
	}

	return nil
}

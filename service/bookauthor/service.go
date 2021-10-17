package bookauthor

import (
	"errors"

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

func (bookauthorService *Service) StoreBatch(bookauthorDomains []bookauthor.Domain) error {
	var err error

	if len(bookauthorDomains) == 0 {
		return errors.New("empty")
	}

	err = bookauthorService.repository.DeleteByBookID(bookauthorDomains[0].BookID)
	if err != nil {
		return err
	}

	for _, bookauthorDomain := range bookauthorDomains {
		_, err = bookauthorService.repository.Store(bookauthorDomain)
		if err != nil {
			return err
		}
	}

	return nil
}

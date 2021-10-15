package author

import (
	"github.com/daniel5u/suisei/domain/author"
)

type Service struct {
	repository author.Repository
}

func NewService(authorRepository author.Repository) author.Service {
	return &Service{
		repository: authorRepository,
	}
}

func (authorService *Service) Fetch() ([]author.Domain, error) {
	result, err := authorService.repository.Fetch()
	if err != nil {
		return []author.Domain{}, err
	}

	return result, nil
}

func (authorService *Service) GetByID(id int) (author.Domain, error) {
	result, err := authorService.repository.GetByID(id)
	if err != nil {
		return author.Domain{}, err
	}

	return result, nil
}

func (authorService *Service) GetByName(name string) (author.Domain, error) {
	result, err := authorService.repository.GetByName(name)
	if err != nil {
		return author.Domain{}, err
	}

	return result, nil
}

func (authorService *Service) Store(authorDomain author.Domain) (author.Domain, error) {
	result, err := authorService.repository.Store(authorDomain)
	if err != nil {
		return author.Domain{}, err
	}

	return result, nil
}

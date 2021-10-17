package publisher

import (
	"github.com/daniel5u/suisei/domain/publisher"
)

type Service struct {
	repository publisher.Repository
}

func NewService(publisherRepository publisher.Repository) publisher.Service {
	return &Service{
		repository: publisherRepository,
	}
}

func (publisherService *Service) Fetch() ([]publisher.Domain, error) {
	result, err := publisherService.repository.Fetch()
	if err != nil {
		return []publisher.Domain{}, err
	}

	return result, nil
}

func (publisherService *Service) GetByID(id int) (publisher.Domain, error) {
	result, err := publisherService.repository.GetByID(id)
	if err != nil {
		return publisher.Domain{}, err
	}

	return result, nil
}

func (publisherService *Service) GetByName(name string) (publisher.Domain, error) {
	result, err := publisherService.repository.GetByName(name)
	if err != nil {
		return publisher.Domain{}, err
	}

	return result, nil
}

func (publisherService *Service) Store(publisherDomain publisher.Domain) (publisher.Domain, error) {
	result, err := publisherService.repository.Store(publisherDomain)
	if err != nil {
		return publisher.Domain{}, err
	}

	return result, nil
}

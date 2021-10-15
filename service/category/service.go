package category

import (
	"github.com/daniel5u/suisei/domain/category"
)

type Service struct {
	repository category.Repository
}

func NewService(categoryRepository category.Repository) category.Service {
	return &Service{
		repository: categoryRepository,
	}
}

func (categoryService *Service) Fetch() ([]category.Domain, error) {
	result, err := categoryService.repository.Fetch()
	if err != nil {
		return []category.Domain{}, err
	}

	return result, nil
}

func (categoryService *Service) GetByID(id int) (category.Domain, error) {
	result, err := categoryService.repository.GetByID(id)
	if err != nil {
		return category.Domain{}, err
	}

	return result, nil
}

func (categoryService *Service) GetByName(name string) (category.Domain, error) {
	result, err := categoryService.repository.GetByName(name)
	if err != nil {
		return category.Domain{}, err
	}

	return result, nil
}

func (categoryService *Service) Store(categoryDomain category.Domain) (category.Domain, error) {
	result, err := categoryService.repository.Store(categoryDomain)
	if err != nil {
		return category.Domain{}, err
	}

	return result, nil
}

package user

import (
	"github.com/daniel5u/suisei/domain/user"
)

type Service struct {
	repository user.Repository
}

func NewService(userRepository user.Repository) user.Service {
	return &Service{
		repository: userRepository,
	}
}

func (userService *Service) Fetch() ([]user.Domain, error) {
	result, err := userService.repository.Fetch()
	if err != nil {
		return []user.Domain{}, err
	}

	return result, nil
}

func (userService *Service) GetByID(id int) (user.Domain, error) {
	result, err := userService.repository.GetByID(id)
	if err != nil {
		return user.Domain{}, err
	}

	return result, nil
}

func (userService *Service) Update(userDomain user.Domain, id int) (user.Domain, error) {
	result, err := userService.repository.Update(userDomain, id)
	if err != nil {
		return user.Domain{}, err
	}

	return result, nil
}

func (userService *Service) Store(userDomain user.Domain) (user.Domain, error) {
	result, err := userService.repository.Store(userDomain)
	if err != nil {
		return user.Domain{}, err
	}

	return result, nil
}

func (userService *Service) Delete(id int) error {
	err := userService.repository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}

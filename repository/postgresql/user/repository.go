package user

import (
	"github.com/daniel5u/suisei/domain/user"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) user.Repository {
	return &Repository{
		DB: db,
	}
}

func (userRepository *Repository) Fetch() ([]user.Domain, error) {
	var userRecords []User
	var userDomains []user.Domain

	err := userRepository.DB.Find(&userRecords).Error
	if err != nil {
		return []user.Domain{}, err
	}

	for _, userRecord := range userRecords {
		userDomains = append(userDomains, repositoryToDomain(userRecord))
	}

	return userDomains, nil
}

func (userRepository *Repository) GetByID(id int) (user.Domain, error) {
	var userRecord User

	err := userRepository.DB.Where("id = ?", id).First(&userRecord).Error
	if err != nil {
		return user.Domain{}, err
	}

	userDomain := repositoryToDomain(userRecord)

	return userDomain, nil
}

func (userRepository *Repository) Update(userDomain user.Domain, id int) (user.Domain, error) {
	var userRecord User = domainToRepository(userDomain)
	var userRecordAfter User

	err := userRepository.DB.Where("id = ?", id).Updates(&userRecord).Error
	if err != nil {
		return user.Domain{}, err
	}

	// get updated row
	err = userRepository.DB.Where("id = ?", id).First(&userRecordAfter).Error
	if err != nil {
		return user.Domain{}, err
	}

	userDomainAfter := repositoryToDomain(userRecordAfter)

	return userDomainAfter, nil
}

func (userRepository *Repository) UpdateBalance(userDomain user.Domain, id int) error {
	var userRecord User = domainToRepository(userDomain)

	err := userRepository.DB.Where("id = ?", id).Select("balance").Updates(&userRecord).Error
	if err != nil {
		return err
	}

	return nil
}

func (userRepository *Repository) Store(userDomain user.Domain) (user.Domain, error) {
	var userRecord User = domainToRepository(userDomain)

	err := userRepository.DB.Create(&userRecord).Error
	if err != nil {
		return user.Domain{}, err
	}

	userDomainAfter := repositoryToDomain(userRecord)

	return userDomainAfter, nil
}

func (userRepository *Repository) Delete(id int) error {
	var userRecord User

	err := userRepository.DB.Where("id = ?", id).Delete(&userRecord).Error
	if err != nil {
		return err
	}

	return nil
}

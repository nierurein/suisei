package category

import (
	"github.com/daniel5u/suisei/domain/category"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) category.Repository {
	return &Repository{
		DB: db,
	}
}

func (categoryRepository *Repository) Fetch() ([]category.Domain, error) {
	var categoryRecords []Category
	var categoryDomains []category.Domain

	err := categoryRepository.DB.Find(&categoryRecords).Error
	if err != nil {
		return []category.Domain{}, err
	}

	for _, categoryRecord := range categoryRecords {
		categoryDomains = append(categoryDomains, repositoryToDomain(categoryRecord))
	}

	return categoryDomains, nil
}

func (categoryRepository *Repository) GetByID(id int) (category.Domain, error) {
	var categoryRecord Category

	err := categoryRepository.DB.Where("id = ?", id).First(&categoryRecord).Error
	if err != nil {
		return category.Domain{}, err
	}

	categoryDomain := repositoryToDomain(categoryRecord)

	return categoryDomain, nil
}

func (categoryRepository *Repository) GetByName(name string) (category.Domain, error) {
	var categoryRecord Category

	err := categoryRepository.DB.Where("id = ?", name).First(&categoryRecord).Error
	if err != nil {
		return category.Domain{}, err
	}

	categoryDomain := repositoryToDomain(categoryRecord)

	return categoryDomain, nil
}

func (categoryRepository *Repository) Store(categoryDomain category.Domain) (category.Domain, error) {
	var categoryRecord Category = domainToRepository(categoryDomain)

	err := categoryRepository.DB.Create(&categoryRecord).Error
	if err != nil {
		return category.Domain{}, err
	}

	categoryDomainAfter := repositoryToDomain(categoryRecord)

	return categoryDomainAfter, nil
}

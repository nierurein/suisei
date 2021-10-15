package author

import (
	"github.com/daniel5u/suisei/domain/author"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) author.Repository {
	return &Repository{
		DB: db,
	}
}

func (authorRepository *Repository) Fetch() ([]author.Domain, error) {
	var authorRecords []Author
	var authorDomains []author.Domain

	err := authorRepository.DB.Find(&authorRecords).Error
	if err != nil {
		return []author.Domain{}, err
	}

	for _, authorRecord := range authorRecords {
		authorDomains = append(authorDomains, repositoryToDomain(authorRecord))
	}

	return authorDomains, nil
}

func (authorRepository *Repository) GetByID(id int) (author.Domain, error) {
	var authorRecord Author

	err := authorRepository.DB.Where("id = ?", id).First(&authorRecord).Error
	if err != nil {
		return author.Domain{}, err
	}

	authorDomain := repositoryToDomain(authorRecord)

	return authorDomain, nil
}

func (authorRepository *Repository) GetByName(name string) (author.Domain, error) {
	var authorRecord Author

	err := authorRepository.DB.Where("id = ?", name).First(&authorRecord).Error
	if err != nil {
		return author.Domain{}, err
	}

	authorDomain := repositoryToDomain(authorRecord)

	return authorDomain, nil
}

func (authorRepository *Repository) Store(authorDomain author.Domain) (author.Domain, error) {
	var authorRecord Author = domainToRepository(authorDomain)

	err := authorRepository.DB.Create(&authorRecord).Error
	if err != nil {
		return author.Domain{}, err
	}

	authorDomainAfter := repositoryToDomain(authorRecord)

	return authorDomainAfter, nil
}

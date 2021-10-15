package publisher

import (
	"github.com/daniel5u/suisei/domain/publisher"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) publisher.Repository {
	return &Repository{
		DB: db,
	}
}

func (publisherRepository *Repository) Fetch() ([]publisher.Domain, error) {
	var publisherRecords []Publisher
	var publisherDomains []publisher.Domain

	err := publisherRepository.DB.Find(&publisherRecords).Error
	if err != nil {
		return []publisher.Domain{}, err
	}

	for _, publisherRecord := range publisherRecords {
		publisherDomains = append(publisherDomains, repositoryToDomain(publisherRecord))
	}

	return publisherDomains, nil
}

func (publisherRepository *Repository) GetByID(id int) (publisher.Domain, error) {
	var publisherRecord Publisher

	err := publisherRepository.DB.Where("id = ?", id).First(&publisherRecord).Error
	if err != nil {
		return publisher.Domain{}, err
	}

	publisherDomain := repositoryToDomain(publisherRecord)

	return publisherDomain, nil
}

func (publisherRepository *Repository) GetByName(name string) (publisher.Domain, error) {
	var publisherRecord Publisher

	err := publisherRepository.DB.Where("id = ?", name).First(&publisherRecord).Error
	if err != nil {
		return publisher.Domain{}, err
	}

	publisherDomain := repositoryToDomain(publisherRecord)

	return publisherDomain, nil
}

func (publisherRepository *Repository) Store(publisherDomain publisher.Domain) (publisher.Domain, error) {
	var publisherRecord Publisher = domainToRepository(publisherDomain)

	err := publisherRepository.DB.Create(&publisherRecord).Error
	if err != nil {
		return publisher.Domain{}, err
	}

	publisherDomainAfter := repositoryToDomain(publisherRecord)

	return publisherDomainAfter, nil
}

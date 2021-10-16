package book

import (
	"github.com/daniel5u/suisei/domain/book"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) book.Repository {
	return &Repository{
		DB: db,
	}
}

func (bookRepository *Repository) Fetch() ([]book.Domain, error) {
	var bookRecords []Book
	var bookDomains []book.Domain

	err := bookRepository.DB.Find(&bookRecords).Error
	if err != nil {
		return []book.Domain{}, err
	}

	for _, bookRecord := range bookRecords {
		bookDomains = append(bookDomains, repositoryToDomain(bookRecord))
	}

	return bookDomains, nil
}

func (bookRepository *Repository) GetByID(id int) (book.Domain, error) {
	var bookRecord Book

	err := bookRepository.DB.Where("id = ?", id).First(&bookRecord).Error
	if err != nil {
		return book.Domain{}, err
	}

	bookDomain := repositoryToDomain(bookRecord)

	return bookDomain, nil
}

func (bookRepository *Repository) Update(bookDomain book.Domain, id int) (book.Domain, error) {
	var bookRecord Book = domainToRepository(bookDomain)
	var bookRecordAfter Book

	err := bookRepository.DB.Where("id = ?", id).Updates(&bookRecord).Error
	if err != nil {
		return book.Domain{}, err
	}

	// get updated row
	err = bookRepository.DB.Where("id = ?", id).First(&bookRecordAfter).Error
	if err != nil {
		return book.Domain{}, err
	}

	bookDomainAfter := repositoryToDomain(bookRecordAfter)

	return bookDomainAfter, nil
}

func (bookRepository *Repository) Store(bookDomain book.Domain) (book.Domain, error) {
	var bookRecord Book = domainToRepository(bookDomain)

	err := bookRepository.DB.Create(&bookRecord).Error
	if err != nil {
		return book.Domain{}, err
	}

	bookDomainAfter := repositoryToDomain(bookRecord)

	return bookDomainAfter, nil
}

func (bookRepository *Repository) Delete(id int) error {
	var bookRecord Book

	err := bookRepository.DB.Where("id = ?", id).Delete(&bookRecord).Error
	if err != nil {
		return err
	}

	return nil
}

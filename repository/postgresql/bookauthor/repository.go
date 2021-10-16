package bookauthor

import (
	"github.com/daniel5u/suisei/domain/bookauthor"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) bookauthor.Repository {
	return &Repository{
		DB: db,
	}
}

func (bookauthorRepository *Repository) Store(bookauthorDomain bookauthor.Domain) (bookauthor.Domain, error) {
	var bookauthorRecord BookAuthor = domainToRepository(bookauthorDomain)

	err := bookauthorRepository.DB.Create(&bookauthorRecord).Error
	if err != nil {
		return bookauthor.Domain{}, err
	}

	bookauthorDomainAfter := repositoryToDomain(bookauthorRecord)

	return bookauthorDomainAfter, nil
}

func (bookauthorRepository *Repository) DeleteByBookID(bookid int) error {
	var bookauthorRecord BookAuthor

	// permanent delete
	err := bookauthorRepository.DB.Unscoped().Where("book_id = ?", bookid).Delete(&bookauthorRecord).Error
	if err != nil {
		return err
	}

	return nil
}

func (bookauthorRepository *Repository) DeleteByAuthorID(authorid int) error {
	var bookauthorRecord BookAuthor

	// permanent delete
	err := bookauthorRepository.DB.Unscoped().Where("author_id = ?", authorid).Delete(&bookauthorRecord).Error
	if err != nil {
		return err
	}

	return nil
}

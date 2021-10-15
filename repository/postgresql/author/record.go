package author

import (
	"github.com/daniel5u/suisei/domain/author"
	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	Name string
}

func repositoryToDomain(authorRepository Author) author.Domain {
	return author.Domain{
		ID:        int(authorRepository.ID),
		Name:      authorRepository.Name,
		CreatedAt: authorRepository.CreatedAt,
		UpdatedAt: authorRepository.UpdatedAt,
	}
}

func domainToRepository(authorDomain author.Domain) Author {
	return Author{
		Name: authorDomain.Name,
	}
}

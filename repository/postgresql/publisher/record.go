package publisher

import (
	"github.com/daniel5u/suisei/domain/publisher"
	"github.com/daniel5u/suisei/repository/postgresql/book"
	"gorm.io/gorm"
)

type Publisher struct {
	gorm.Model
	Name  string
	Books []book.Book
}

func repositoryToDomain(publisherRepository Publisher) publisher.Domain {
	return publisher.Domain{
		ID:        int(publisherRepository.ID),
		Name:      publisherRepository.Name,
		CreatedAt: publisherRepository.CreatedAt,
		UpdatedAt: publisherRepository.UpdatedAt,
	}
}

func domainToRepository(publisherDomain publisher.Domain) Publisher {
	return Publisher{
		Name: publisherDomain.Name,
	}
}

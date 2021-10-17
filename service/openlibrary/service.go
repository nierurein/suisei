package openlibrary

import (
	"github.com/daniel5u/suisei/domain/author"
	"github.com/daniel5u/suisei/domain/book"
	"github.com/daniel5u/suisei/domain/bookauthor"
	"github.com/daniel5u/suisei/domain/category"
	"github.com/daniel5u/suisei/domain/openlibrary"
	"github.com/daniel5u/suisei/domain/publisher"
)

type Service struct {
	repository        openlibrary.Repository
	bookService       book.Service
	authorService     author.Service
	bookauthorService bookauthor.Service
	categoryService   category.Service
	publihserService  publisher.Service
}

func NewService(openlibraryRepository openlibrary.Repository, bs book.Service, as author.Service, bas bookauthor.Service, cs category.Service, ps publisher.Service) openlibrary.Service {
	return &Service{
		repository:        openlibraryRepository,
		bookService:       bs,
		authorService:     as,
		bookauthorService: bas,
		categoryService:   cs,
		publihserService:  ps,
	}
}

func (openlibraryRepository *Service) Fetch(links []string) error {
	var openlibraryDomains []openlibrary.Domain
	var categoryDomain category.Domain
	var publisherDomain publisher.Domain
	var bookDomain book.Domain
	var bookDomainAfter book.Domain
	var authorDomain author.Domain
	var bookauthorDomains []bookauthor.Domain
	var err error

	openlibraryDomains, err = openlibraryRepository.repository.Fetch(links)
	if err != nil {
		return err
	}

	for _, openlibraryDomain := range openlibraryDomains {
		categoryDomain, _ = openlibraryRepository.categoryService.GetByName(openlibraryDomain.Category)
		if categoryDomain.Name == "" {
			categoryDomain, err = openlibraryRepository.categoryService.Store(category.Domain{
				Name: openlibraryDomain.Category,
			})
			if err != nil {
				return err
			}
		}

		publisherDomain, _ = openlibraryRepository.publihserService.GetByName(openlibraryDomain.Publisher)
		if publisherDomain.Name == "" {
			publisherDomain, err = openlibraryRepository.publihserService.Store(publisher.Domain{
				Name: openlibraryDomain.Publisher,
			})
			if err != nil {
				return err
			}
		}

		bookDomain.PublisherID = publisherDomain.ID
		bookDomain.CategoryID = categoryDomain.ID
		bookDomain.ISBN = openlibraryDomain.ISBN
		bookDomain.Title = openlibraryDomain.Title
		bookDomain.Price = openlibraryDomain.Price
		bookDomain.PublicationYear = openlibraryDomain.PublicationYear
		bookDomain.PageCount = openlibraryDomain.PageCount
		bookDomain.Description = openlibraryDomain.Description

		bookDomainAfter, err = openlibraryRepository.bookService.Store(bookDomain)
		if err != nil {
			return err
		}

		for _, openlibraryDomainAuthor := range openlibraryDomain.Authors {
			authorDomain, _ = openlibraryRepository.authorService.GetByName(openlibraryDomainAuthor)
			if authorDomain.Name == "" {
				authorDomain, err = openlibraryRepository.authorService.Store(author.Domain{
					Name: openlibraryDomainAuthor,
				})
				if err != nil {
					return err
				}
			}

			bookauthorDomains = append(bookauthorDomains, bookauthor.Domain{
				BookID:   bookDomainAfter.ID,
				AuthorID: authorDomain.ID,
			})
		}

		err = openlibraryRepository.bookauthorService.StoreBatch(bookauthorDomains)
		if err != nil {
			return err
		}

		// clear bookauthorDomains
		bookauthorDomains = nil
	}

	return nil
}

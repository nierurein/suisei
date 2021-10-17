package bookauthor_test

import (
	"errors"
	"os"
	"testing"

	_bookauthorDomain "github.com/daniel5u/suisei/domain/bookauthor"
	_bookauthorMock "github.com/daniel5u/suisei/domain/bookauthor/mocks"
	_bookauthorService "github.com/daniel5u/suisei/service/bookauthor"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	bookauthorRepository _bookauthorMock.Repository
	bookauthorService    _bookauthorDomain.Service
)

func TestMain(m *testing.M) {
	bookauthorService = _bookauthorService.NewService(&bookauthorRepository)
	os.Exit(m.Run())
}

func TestStoreBatch(t *testing.T) {
	t.Run("empty argument", func(t *testing.T) {
		err := bookauthorService.StoreBatch([]_bookauthorDomain.Domain{})

		assert.NotNil(t, err)
	})

	t.Run("valid domain", func(t *testing.T) {
		bookauthorDomainTest := []_bookauthorDomain.Domain{
			{
				BookID:   1,
				AuthorID: 1,
			},
			{
				BookID:   2,
				AuthorID: 2,
			},
		}
		bookauthorRepository.On("DeleteByBookID", mock.Anything).Return(nil).Once()
		bookauthorRepository.On("Store", mock.Anything).Return(bookauthorDomainTest[0], nil, bookauthorDomainTest[1], nil).Twice()

		err := bookauthorService.StoreBatch(bookauthorDomainTest)

		assert.Nil(t, err)
	})

	t.Run("invalid delete", func(t *testing.T) {
		bookauthorDomainTest := []_bookauthorDomain.Domain{
			{
				BookID:   1,
				AuthorID: 1,
			},
			{
				BookID:   2,
				AuthorID: 2,
			},
		}
		bookauthorRepository.On("DeleteByBookID", mock.Anything).Return(errors.New("Record not found")).Once()

		err := bookauthorService.StoreBatch(bookauthorDomainTest)

		assert.NotNil(t, err)
	})

	t.Run("repository error", func(t *testing.T) {
		bookauthorDomainTest := []_bookauthorDomain.Domain{
			{
				BookID:   1,
				AuthorID: 1,
			},
			{
				BookID:   2,
				AuthorID: 2,
			},
		}
		bookauthorRepository.On("DeleteByBookID", mock.Anything).Return(nil).Once()
		bookauthorRepository.On("Store", mock.Anything).Return(_bookauthorDomain.Domain{}, errors.New("Undefined behaviour")).Once()

		err := bookauthorService.StoreBatch(bookauthorDomainTest)

		assert.NotNil(t, err)
	})
}

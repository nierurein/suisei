package book_test

import (
	"errors"
	"os"
	"testing"

	_bookDomain "github.com/daniel5u/suisei/domain/book"
	_bookMock "github.com/daniel5u/suisei/domain/book/mocks"
	_bookService "github.com/daniel5u/suisei/service/book"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	bookRepository _bookMock.Repository
	bookService    _bookDomain.Service
)

func TestMain(m *testing.M) {
	bookService = _bookService.NewService(&bookRepository)
	os.Exit(m.Run())
}

func TestFetch(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		bookDomainTest := []_bookDomain.Domain{
			{
				ID:    1,
				Title: "book",
			},
			{
				ID:    2,
				Title: "book2",
			},
		}
		bookRepository.On("Fetch").Return(bookDomainTest, nil).Once()

		result, err := bookService.Fetch()

		assert.Nil(t, err)
		assert.Equal(t, bookDomainTest, result)
	})

	t.Run("repository error", func(t *testing.T) {
		bookRepository.On("Fetch").Return([]_bookDomain.Domain{}, errors.New("Record not found")).Once()

		result, err := bookService.Fetch()

		assert.NotNil(t, err)
		assert.Empty(t, result)
	})
}

func TestGetByID(t *testing.T) {
	t.Run("valid id", func(t *testing.T) {
		bookDomainTest := _bookDomain.Domain{
			ID:    1,
			Title: "book",
		}
		bookRepository.On("GetByID", mock.AnythingOfType("int")).Return(bookDomainTest, nil).Once()

		result, err := bookService.GetByID(1)

		assert.Nil(t, err)
		assert.Equal(t, bookDomainTest.Title, result.Title)
	})

	t.Run("repository error", func(t *testing.T) {
		bookRepository.On("GetByID", mock.AnythingOfType("int")).Return(_bookDomain.Domain{}, errors.New("Record not found")).Once()

		result, err := bookService.GetByID(-1)

		assert.NotNil(t, err)
		assert.Empty(t, result)
	})
}

func TestGetByTitle(t *testing.T) {
	t.Run("valid name", func(t *testing.T) {
		bookDomainTest := _bookDomain.Domain{
			ID:    1,
			Title: "book",
		}
		bookRepository.On("GetByTitle", mock.AnythingOfType("string")).Return(bookDomainTest, nil).Once()

		result, err := bookService.GetByTitle("book")

		assert.Nil(t, err)
		assert.Equal(t, bookDomainTest.ID, result.ID)
	})

	t.Run("repository error", func(t *testing.T) {
		bookRepository.On("GetByTitle", mock.AnythingOfType("string")).Return(_bookDomain.Domain{}, errors.New("Record not found")).Once()

		result, err := bookService.GetByTitle("thisTitleDoesNotExist")

		assert.NotNil(t, err)
		assert.Empty(t, result)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("valid domain", func(t *testing.T) {
		bookDomainTest := _bookDomain.Domain{
			ID:    1,
			Title: "newbook",
			Price: 100,
		}
		bookDomainUpdate := _bookDomain.Domain{
			Price: 100,
		}
		bookRepository.On("Update", mock.Anything, mock.AnythingOfType("int")).Return(bookDomainTest, nil).Once()

		result, err := bookService.Update(bookDomainUpdate, 1)

		assert.Nil(t, err)
		assert.Equal(t, bookDomainTest, result)
	})

	t.Run("repository error", func(t *testing.T) {
		bookRepository.On("Update", mock.Anything, mock.AnythingOfType("int")).Return(_bookDomain.Domain{}, errors.New("Record not found")).Once()

		result, err := bookService.Update(_bookDomain.Domain{
			Title: "book1",
		}, 1)

		assert.NotNil(t, err)
		assert.Empty(t, result)
	})
}

func TestStore(t *testing.T) {
	t.Run("valid domain", func(t *testing.T) {
		bookDomainTest := _bookDomain.Domain{
			ID:    1,
			Title: "book",
		}
		bookRepository.On("Store", mock.Anything).Return(bookDomainTest, nil).Once()

		result, err := bookService.Store(bookDomainTest)

		assert.Nil(t, err)
		assert.Equal(t, bookDomainTest, result)
	})

	t.Run("repository error", func(t *testing.T) {
		bookRepository.On("Store", mock.Anything).Return(_bookDomain.Domain{}, errors.New("Undefined behaviour")).Once()

		result, err := bookService.Store(_bookDomain.Domain{
			Title: "book1",
		})

		assert.NotNil(t, err)
		assert.Empty(t, result)
	})
}

func TestDelete(t *testing.T) {
	t.Run("valid id", func(t *testing.T) {
		bookRepository.On("Delete", mock.AnythingOfType("int")).Return(nil).Once()

		err := bookService.Delete(1)

		assert.Nil(t, err)
	})

	t.Run("repository error", func(t *testing.T) {
		bookRepository.On("Delete", mock.AnythingOfType("int")).Return(errors.New("Record not found")).Once()

		err := bookService.Delete(-1)

		assert.NotNil(t, err)
	})
}

package author_test

import (
	"errors"
	"os"
	"testing"

	_authorDomain "github.com/daniel5u/suisei/domain/author"
	_authorMock "github.com/daniel5u/suisei/domain/author/mocks"
	_authorService "github.com/daniel5u/suisei/service/author"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	authorRepository _authorMock.Repository
	authorService    _authorDomain.Service
)

func TestMain(m *testing.M) {
	authorService = _authorService.NewService(&authorRepository)
	os.Exit(m.Run())
}

func TestFetch(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		authorDomainTest := []_authorDomain.Domain{
			{
				ID:   1,
				Name: "author",
			},
			{
				ID:   2,
				Name: "author2",
			},
		}
		authorRepository.On("Fetch").Return(authorDomainTest, nil).Once()

		result, err := authorService.Fetch()

		assert.Nil(t, err)
		assert.Equal(t, authorDomainTest, result)
	})

	t.Run("repository error", func(t *testing.T) {
		authorRepository.On("Fetch").Return([]_authorDomain.Domain{}, errors.New("Record not found")).Once()

		result, err := authorService.Fetch()

		assert.NotNil(t, err)
		assert.Empty(t, result)
	})
}

func TestGetByID(t *testing.T) {
	t.Run("valid id", func(t *testing.T) {
		authorDomainTest := _authorDomain.Domain{
			ID:   1,
			Name: "author",
		}
		authorRepository.On("GetByID", mock.AnythingOfType("int")).Return(authorDomainTest, nil).Once()

		result, err := authorService.GetByID(1)

		assert.Nil(t, err)
		assert.Equal(t, authorDomainTest.Name, result.Name)
	})

	t.Run("repository error", func(t *testing.T) {
		authorRepository.On("GetByID", mock.AnythingOfType("int")).Return(_authorDomain.Domain{}, errors.New("Record not found")).Once()

		result, err := authorService.GetByID(-1)

		assert.NotNil(t, err)
		assert.Empty(t, result)
	})
}

func TestGetByName(t *testing.T) {
	t.Run("valid name", func(t *testing.T) {
		authorDomainTest := _authorDomain.Domain{
			ID:   1,
			Name: "author",
		}
		authorRepository.On("GetByName", mock.AnythingOfType("string")).Return(authorDomainTest, nil).Once()

		result, err := authorService.GetByName("author")

		assert.Nil(t, err)
		assert.Equal(t, authorDomainTest.ID, result.ID)
	})

	t.Run("repository error", func(t *testing.T) {
		authorRepository.On("GetByName", mock.AnythingOfType("string")).Return(_authorDomain.Domain{}, errors.New("Record not found")).Once()

		result, err := authorService.GetByName("thisNameDoesNotExist")

		assert.NotNil(t, err)
		assert.Empty(t, result)
	})
}

func TestStore(t *testing.T) {
	t.Run("valid domain", func(t *testing.T) {
		authorDomainTest := _authorDomain.Domain{
			ID:   1,
			Name: "author",
		}
		authorRepository.On("Store", mock.Anything).Return(authorDomainTest, nil).Once()

		result, err := authorService.Store(authorDomainTest)

		assert.Nil(t, err)
		assert.Equal(t, authorDomainTest, result)
	})

	t.Run("repository error", func(t *testing.T) {
		authorRepository.On("Store", mock.Anything).Return(_authorDomain.Domain{}, errors.New("Undefined behaviour")).Once()

		result, err := authorService.Store(_authorDomain.Domain{
			Name: "author1",
		})

		assert.NotNil(t, err)
		assert.Empty(t, result)
	})
}

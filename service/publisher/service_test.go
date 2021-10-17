package publisher_test

import (
	"errors"
	"os"
	"testing"

	_publisherDomain "github.com/daniel5u/suisei/domain/publisher"
	_publisherMock "github.com/daniel5u/suisei/domain/publisher/mocks"
	_publisherService "github.com/daniel5u/suisei/service/publisher"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	publisherRepository _publisherMock.Repository
	publisherService    _publisherDomain.Service
)

func TestMain(m *testing.M) {
	publisherService = _publisherService.NewService(&publisherRepository)
	os.Exit(m.Run())
}

func TestFetch(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		publisherDomainTest := []_publisherDomain.Domain{
			{
				ID:   1,
				Name: "publisher",
			},
			{
				ID:   2,
				Name: "publisher2",
			},
		}
		publisherRepository.On("Fetch").Return(publisherDomainTest, nil).Once()

		result, err := publisherService.Fetch()

		assert.Nil(t, err)
		assert.Equal(t, publisherDomainTest, result)
	})

	t.Run("repository error", func(t *testing.T) {
		publisherRepository.On("Fetch").Return([]_publisherDomain.Domain{}, errors.New("Record not found")).Once()

		result, err := publisherService.Fetch()

		assert.NotNil(t, err)
		assert.Empty(t, result)
	})
}

func TestGetByID(t *testing.T) {
	t.Run("valid id", func(t *testing.T) {
		publisherDomainTest := _publisherDomain.Domain{
			ID:   1,
			Name: "publisher",
		}
		publisherRepository.On("GetByID", mock.AnythingOfType("int")).Return(publisherDomainTest, nil).Once()

		result, err := publisherService.GetByID(1)

		assert.Nil(t, err)
		assert.Equal(t, publisherDomainTest.Name, result.Name)
	})

	t.Run("repository error", func(t *testing.T) {
		publisherRepository.On("GetByID", mock.AnythingOfType("int")).Return(_publisherDomain.Domain{}, errors.New("Record not found")).Once()

		result, err := publisherService.GetByID(-1)

		assert.NotNil(t, err)
		assert.Empty(t, result)
	})
}

func TestGetByName(t *testing.T) {
	t.Run("valid name", func(t *testing.T) {
		publisherDomainTest := _publisherDomain.Domain{
			ID:   1,
			Name: "publisher",
		}
		publisherRepository.On("GetByName", mock.AnythingOfType("string")).Return(publisherDomainTest, nil).Once()

		result, err := publisherService.GetByName("publisher")

		assert.Nil(t, err)
		assert.Equal(t, publisherDomainTest.ID, result.ID)
	})

	t.Run("repository error", func(t *testing.T) {
		publisherRepository.On("GetByName", mock.AnythingOfType("string")).Return(_publisherDomain.Domain{}, errors.New("Record not found")).Once()

		result, err := publisherService.GetByName("thisNameDoesNotExist")

		assert.NotNil(t, err)
		assert.Empty(t, result)
	})
}

func TestStore(t *testing.T) {
	t.Run("valid domain", func(t *testing.T) {
		publisherDomainTest := _publisherDomain.Domain{
			ID:   1,
			Name: "publisher",
		}
		publisherRepository.On("Store", mock.Anything).Return(publisherDomainTest, nil).Once()

		result, err := publisherService.Store(publisherDomainTest)

		assert.Nil(t, err)
		assert.Equal(t, publisherDomainTest, result)
	})

	t.Run("repository error", func(t *testing.T) {
		publisherRepository.On("Store", mock.Anything).Return(_publisherDomain.Domain{}, errors.New("Undefined behaviour")).Once()

		result, err := publisherService.Store(_publisherDomain.Domain{
			Name: "publisher1",
		})

		assert.NotNil(t, err)
		assert.Empty(t, result)
	})
}

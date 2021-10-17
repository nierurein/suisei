package category_test

import (
	"errors"
	"os"
	"testing"

	_categoryDomain "github.com/daniel5u/suisei/domain/category"
	_categoryMock "github.com/daniel5u/suisei/domain/category/mocks"
	_categoryService "github.com/daniel5u/suisei/service/category"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	categoryRepository _categoryMock.Repository
	categoryService    _categoryDomain.Service
)

func TestMain(m *testing.M) {
	categoryService = _categoryService.NewService(&categoryRepository)
	os.Exit(m.Run())
}

func TestFetch(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		categoryDomainTest := []_categoryDomain.Domain{
			{
				ID:   1,
				Name: "category",
			},
			{
				ID:   2,
				Name: "category2",
			},
		}
		categoryRepository.On("Fetch").Return(categoryDomainTest, nil).Once()

		result, err := categoryService.Fetch()

		assert.Nil(t, err)
		assert.Equal(t, categoryDomainTest, result)
	})

	t.Run("repository error", func(t *testing.T) {
		categoryRepository.On("Fetch").Return([]_categoryDomain.Domain{}, errors.New("Record not found")).Once()

		result, err := categoryService.Fetch()

		assert.NotNil(t, err)
		assert.Empty(t, result)
	})
}

func TestGetByID(t *testing.T) {
	t.Run("valid id", func(t *testing.T) {
		categoryDomainTest := _categoryDomain.Domain{
			ID:   1,
			Name: "category",
		}
		categoryRepository.On("GetByID", mock.AnythingOfType("int")).Return(categoryDomainTest, nil).Once()

		result, err := categoryService.GetByID(1)

		assert.Nil(t, err)
		assert.Equal(t, categoryDomainTest.Name, result.Name)
	})

	t.Run("repository error", func(t *testing.T) {
		categoryRepository.On("GetByID", mock.AnythingOfType("int")).Return(_categoryDomain.Domain{}, errors.New("Record not found")).Once()

		result, err := categoryService.GetByID(-1)

		assert.NotNil(t, err)
		assert.Empty(t, result)
	})
}

func TestGetByName(t *testing.T) {
	t.Run("valid name", func(t *testing.T) {
		categoryDomainTest := _categoryDomain.Domain{
			ID:   1,
			Name: "category",
		}
		categoryRepository.On("GetByName", mock.AnythingOfType("string")).Return(categoryDomainTest, nil).Once()

		result, err := categoryService.GetByName("category")

		assert.Nil(t, err)
		assert.Equal(t, categoryDomainTest.ID, result.ID)
	})

	t.Run("repository error", func(t *testing.T) {
		categoryRepository.On("GetByName", mock.AnythingOfType("string")).Return(_categoryDomain.Domain{}, errors.New("Record not found")).Once()

		result, err := categoryService.GetByName("thisNameDoesNotExist")

		assert.NotNil(t, err)
		assert.Empty(t, result)
	})
}

func TestStore(t *testing.T) {
	t.Run("valid domain", func(t *testing.T) {
		categoryDomainTest := _categoryDomain.Domain{
			ID:   1,
			Name: "category",
		}
		categoryRepository.On("Store", mock.Anything).Return(categoryDomainTest, nil).Once()

		result, err := categoryService.Store(categoryDomainTest)

		assert.Nil(t, err)
		assert.Equal(t, categoryDomainTest, result)
	})

	t.Run("repository error", func(t *testing.T) {
		categoryRepository.On("Store", mock.Anything).Return(_categoryDomain.Domain{}, errors.New("Undefined behaviour")).Once()

		result, err := categoryService.Store(_categoryDomain.Domain{
			Name: "category1",
		})

		assert.NotNil(t, err)
		assert.Empty(t, result)
	})
}

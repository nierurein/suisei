package user_test

import (
	"errors"
	"os"
	"testing"

	_userDomain "github.com/daniel5u/suisei/domain/user"
	_userMock "github.com/daniel5u/suisei/domain/user/mocks"
	_userService "github.com/daniel5u/suisei/service/user"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	userRepository _userMock.Repository
	userService    _userDomain.Service
)

func TestMain(m *testing.M) {
	userService = _userService.NewService(&userRepository)
	os.Exit(m.Run())
}

func TestFetch(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		userDomainTest := []_userDomain.Domain{
			{
				ID:   1,
				Name: "user",
			},
			{
				ID:   2,
				Name: "user2",
			},
		}
		userRepository.On("Fetch").Return(userDomainTest, nil).Once()

		result, err := userService.Fetch()

		assert.Nil(t, err)
		assert.Equal(t, userDomainTest, result)
	})

	t.Run("repository error", func(t *testing.T) {
		userRepository.On("Fetch").Return([]_userDomain.Domain{}, errors.New("Record not found")).Once()

		result, err := userService.Fetch()

		assert.NotNil(t, err)
		assert.Empty(t, result)
	})
}

func TestGetByID(t *testing.T) {
	t.Run("valid id", func(t *testing.T) {
		userDomainTest := _userDomain.Domain{
			ID:   1,
			Name: "user",
		}
		userRepository.On("GetByID", mock.AnythingOfType("int")).Return(userDomainTest, nil).Once()

		result, err := userService.GetByID(1)

		assert.Nil(t, err)
		assert.Equal(t, userDomainTest.Name, result.Name)
	})

	t.Run("repository error", func(t *testing.T) {
		userRepository.On("GetByID", mock.AnythingOfType("int")).Return(_userDomain.Domain{}, errors.New("Record not found")).Once()

		result, err := userService.GetByID(-1)

		assert.NotNil(t, err)
		assert.Empty(t, result)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("valid domain", func(t *testing.T) {
		userDomainTest := _userDomain.Domain{
			ID:      1,
			Name:    "newuser",
			Balance: 100,
		}
		userDomainUpdate := _userDomain.Domain{
			Name: "newuser",
		}
		userRepository.On("Update", mock.Anything, mock.AnythingOfType("int")).Return(userDomainTest, nil).Once()

		result, err := userService.Update(userDomainUpdate, 1)

		assert.Nil(t, err)
		assert.Equal(t, userDomainTest, result)
	})

	t.Run("repository error", func(t *testing.T) {
		userRepository.On("Update", mock.Anything, mock.AnythingOfType("int")).Return(_userDomain.Domain{}, errors.New("Record not found")).Once()

		result, err := userService.Update(_userDomain.Domain{
			Name: "user1",
		}, 1)

		assert.NotNil(t, err)
		assert.Empty(t, result)
	})
}

func TestUpdateBalance(t *testing.T) {
	t.Run("valid domain", func(t *testing.T) {
		userDomainUpdate := _userDomain.Domain{
			Balance: 0,
		}
		userRepository.On("UpdateBalance", mock.Anything, mock.AnythingOfType("int")).Return(nil).Once()

		err := userService.UpdateBalance(userDomainUpdate, 1)

		assert.Nil(t, err)
	})

	t.Run("repository error", func(t *testing.T) {
		userRepository.On("UpdateBalance", mock.Anything, mock.AnythingOfType("int")).Return(errors.New("Record not found")).Once()
		err := userService.UpdateBalance(_userDomain.Domain{
			Balance: 1,
		}, 1)

		assert.NotNil(t, err)
	})
}

func TestStore(t *testing.T) {
	t.Run("valid domain", func(t *testing.T) {
		userDomainTest := _userDomain.Domain{
			ID:   1,
			Name: "user",
		}
		userRepository.On("Store", mock.Anything).Return(userDomainTest, nil).Once()

		result, err := userService.Store(userDomainTest)

		assert.Nil(t, err)
		assert.Equal(t, userDomainTest, result)
	})

	t.Run("repository error", func(t *testing.T) {
		userRepository.On("Store", mock.Anything).Return(_userDomain.Domain{}, errors.New("Undefined behaviour")).Once()

		result, err := userService.Store(_userDomain.Domain{
			Name: "user1",
		})

		assert.NotNil(t, err)
		assert.Empty(t, result)
	})
}

func TestDelete(t *testing.T) {
	t.Run("valid id", func(t *testing.T) {
		userRepository.On("Delete", mock.AnythingOfType("int")).Return(nil).Once()

		err := userService.Delete(1)

		assert.Nil(t, err)
	})

	t.Run("repository error", func(t *testing.T) {
		userRepository.On("Delete", mock.AnythingOfType("int")).Return(errors.New("Record not found")).Once()

		err := userService.Delete(-1)

		assert.NotNil(t, err)
	})
}

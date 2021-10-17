package user

import (
	"github.com/daniel5u/suisei/domain/user"
	"github.com/daniel5u/suisei/repository/postgresql/transaction"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name         string
	Email        string
	Password     string
	Balance      int
	Token        string
	Transactions []transaction.Transaction
}

func repositoryToDomain(userRepository User) user.Domain {
	return user.Domain{
		ID:        int(userRepository.ID),
		Name:      userRepository.Name,
		Email:     userRepository.Email,
		Password:  userRepository.Password,
		Balance:   userRepository.Balance,
		Token:     userRepository.Token,
		CreatedAt: userRepository.CreatedAt,
		UpdatedAt: userRepository.UpdatedAt,
	}
}

func domainToRepository(userDomain user.Domain) User {
	return User{
		Name:     userDomain.Name,
		Email:    userDomain.Email,
		Password: userDomain.Password,
		Balance:  userDomain.Balance,
		Token:    userDomain.Token,
	}
}

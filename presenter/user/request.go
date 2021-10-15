package user

import "github.com/daniel5u/suisei/domain/user"

type Request struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Balance  int    `json:"balance"`
	Token    string `json:"token"`
}

func requestToDomain(request Request) user.Domain {
	return user.Domain{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
		Balance:  request.Balance,
		Token:    request.Token,
	}
}

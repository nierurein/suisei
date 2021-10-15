package user

import (
	"time"

	"github.com/daniel5u/suisei/domain/user"
)

type Response struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Balance   int       `json:"balance"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func domainToResponse(userDomain user.Domain) Response {
	return Response{
		ID:        userDomain.ID,
		Name:      userDomain.Name,
		Email:     userDomain.Email,
		Balance:   userDomain.Balance,
		Token:     userDomain.Token,
		CreatedAt: userDomain.CreatedAt,
		UpdatedAt: userDomain.UpdatedAt,
	}
}

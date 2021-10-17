package author

import (
	"time"

	"github.com/daniel5u/suisei/domain/author"
)

type Response struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func domainToResponse(authorDomain author.Domain) Response {
	return Response{
		ID:        authorDomain.ID,
		Name:      authorDomain.Name,
		CreatedAt: authorDomain.CreatedAt,
		UpdatedAt: authorDomain.UpdatedAt,
	}
}

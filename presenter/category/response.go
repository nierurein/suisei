package category

import (
	"time"

	"github.com/daniel5u/suisei/domain/category"
)

type Response struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func domainToResponse(categoryDomain category.Domain) Response {
	return Response{
		ID:        categoryDomain.ID,
		Name:      categoryDomain.Name,
		CreatedAt: categoryDomain.CreatedAt,
		UpdatedAt: categoryDomain.UpdatedAt,
	}
}

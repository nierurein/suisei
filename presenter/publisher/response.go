package publisher

import (
	"time"

	"github.com/daniel5u/suisei/domain/publisher"
)

type Response struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func domainToResponse(publisherDomain publisher.Domain) Response {
	return Response{
		ID:        publisherDomain.ID,
		Name:      publisherDomain.Name,
		CreatedAt: publisherDomain.CreatedAt,
		UpdatedAt: publisherDomain.UpdatedAt,
	}
}

package transaction

import (
	"time"

	"github.com/daniel5u/suisei/domain/transaction"
)

type Response struct {
	ID         int       `json:"id"`
	UserID     int       `json:"user_id"`
	TotalPrice int       `json:"total_price"`
	Status     int       `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func domainToResponse(transactionDomain transaction.Domain) Response {
	return Response{
		ID:         transactionDomain.ID,
		UserID:     transactionDomain.UserID,
		TotalPrice: transactionDomain.TotalPrice,
		Status:     transactionDomain.Status,
		CreatedAt:  transactionDomain.CreatedAt,
		UpdatedAt:  transactionDomain.UpdatedAt,
	}
}

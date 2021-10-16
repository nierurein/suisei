package transaction

import "github.com/daniel5u/suisei/domain/transaction"

type Request struct {
	UserID int `json:"user_id"`
	Status int `json:"status"`
}

func requestToDomain(request Request) transaction.Domain {
	return transaction.Domain{
		UserID: request.UserID,
		Status: request.Status,
	}
}

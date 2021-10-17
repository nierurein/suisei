package booktransaction

import "github.com/daniel5u/suisei/domain/booktransaction"

type RequestItem struct {
	BookID   int `json:"book_id"`
	Quantity int `json:"quantity"`
}

type Request struct {
	Items []RequestItem `json:"items"`
}

func requestItemToDomain(requestItem RequestItem) booktransaction.Domain {
	return booktransaction.Domain{
		BookID:   requestItem.BookID,
		Quantity: requestItem.Quantity,
	}
}

package bookauthor

import "github.com/daniel5u/suisei/domain/bookauthor"

type RequestItem struct {
	AuthorID int `json:"author_id"`
	BookID   int `json:"book_id"`
}

type Request struct {
	Items []RequestItem `json:"items"`
}

func requestItemToDomain(requestItem RequestItem) bookauthor.Domain {
	return bookauthor.Domain{
		AuthorID: requestItem.AuthorID,
		BookID:   requestItem.BookID,
	}
}

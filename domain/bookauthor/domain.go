package bookauthor

type Domain struct {
	BookID   int
	AuthorID int
}

type Service interface {
	StoreBatch(bookauthorDomain []Domain) error
}

type Repository interface {
	Store(bookauthorDomain Domain) (Domain, error)
	DeleteByBookID(bookid int) error
}

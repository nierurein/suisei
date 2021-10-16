package bookauthor

type Domain struct {
	BookID   int
	AuthorID int
}

type Service interface {
	Store(bookauthorDomain Domain) (Domain, error)
	DeleteByBookID(bookid int) error
	DeleteByAuthorID(authorid int) error
}

type Repository interface {
	Store(bookauthorDomain Domain) (Domain, error)
	DeleteByBookID(bookid int) error
	DeleteByAuthorID(authorid int) error
}

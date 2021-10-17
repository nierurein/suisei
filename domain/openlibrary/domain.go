package openlibrary

type Domain struct {
	Publisher       string
	Category        string
	ISBN            string
	Authors         []string
	Title           string
	Price           int
	PublicationYear int
	PageCount       int
	Description     string
}

type Service interface {
	Fetch(links []string) error
}

type Repository interface {
	Fetch(links []string) ([]Domain, error)
}

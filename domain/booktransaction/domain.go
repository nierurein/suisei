package booktransaction

type Domain struct {
	TransactionID int
	BookID        int
	Quantity      int
	PricePerUnit  int
}

type Service interface {
	StoreBatch(booktransactionDomain []Domain, transactionid int) error
}

type Repository interface {
	Store(booktransactionDomain Domain) (Domain, error)
	DeleteByTransactionID(transactionid int) error
}

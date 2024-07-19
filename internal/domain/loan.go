package domain

type Loan struct {
	ID       int
	Book     Book
	Person   Person
	DueDate  string
	Returned bool
}

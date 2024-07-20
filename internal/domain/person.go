package domain

type Address struct {
	City string
}

type Person struct {
	ID      int
	Name    string
	Address Address
}

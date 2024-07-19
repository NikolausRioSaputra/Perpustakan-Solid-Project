package domain

type Address struct {
	Kota string
}

type Person struct {
	ID      int
	Name    string
	Address Address
}

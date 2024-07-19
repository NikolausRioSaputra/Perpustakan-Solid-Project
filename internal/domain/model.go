package domain

type Alamat struct {
	Kota string
}

type Orang struct {
	Nama string
	Alamat
	Buku []string
}

type Buku struct {
	Judul string
	Stock int
}

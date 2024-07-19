package repository

import (
	"Project_Niko/internal/domain"
	"errors"
	"fmt"
)

type BookRepositoryInterface interface {
	BookSaver
	BookUpdater
	BookDeleter
}

// ini interface implementasi di repo untuk nyimpan buku
type BookSaver interface {
	SaveBook(book *domain.Book) error
}

// ini interface untuk implemetasi d repo untuk update buku
type BookUpdater interface {
	UpdateBook(book *domain.Book) error
}

// ini interface untuk implementasi d repo untuk hapus buku
type BookDeleter interface {
	DeleteBook(bookID int) error
}

// ini kita buat untuk nyimpan fungsi bookrepositoryinterface dengan map yang di simpan di dengan books
type BookRepository struct {
	books map[int]domain.Book
}

// Fungsi ini mengembalikan pointer ke sebuah instance BookRepository yang baru.
// Ini memastikan bahwa fungsi ini mengembalikan referensi ke struct BookRepository,
func NewBookRepository() BookRepositoryInterface {
	return &BookRepository{
		//  Map books diinisialisasi sebagai map kosong. ini adalah struktur data yang akan digunakan untuk menyimpan buku-buku dengan ID sebagai kunci dan objek buku sebagai nilai.
		books: map[int]domain.Book{},
	}
}

func (repo *BookRepository) SaveBook(book *domain.Book) error {
	if _, exists := repo.books[book.ID]; exists {
		return errors.New("book already exist")
	}

	repo.books[book.ID] = *book
	fmt.Println(repo.books)
	return nil
}

func (repo *BookRepository) UpdateBook(book *domain.Book) error {
	if _, exists := repo.books[book.ID]; exists {
		repo.books[book.ID] = *book
		return nil
	}

	return errors.New("book not found")
}

func (repo *BookRepository) DeleteBook(bookID int) error {
	if _, exists := repo.books[bookID]; !exists {
		return errors.New("book not found")
	}

	delete(repo.books, bookID)
	return nil
}

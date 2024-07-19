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
	BookReader
}

type BookSaver interface {
	SaveBook(book *domain.Buku) error
}

type BookUpdater interface {
	UpdateBook(book *domain.Buku) error
}

type BookDeleter interface {
	DeleteBook(bookID int) error
}

type BookReader interface {
	GetAllBooks() ([]domain.Buku, error)
	GetBookById(bookID int) (*domain.Buku, error)
}

type BookRepository struct {
	books map[int]domain.Buku
}

func NewBookRepository() BookRepositoryInterface {
	return &BookRepository{
		books: map[int]domain.Buku{},
	}
}

// DeleteBook implements BookRepositoryInterface.
func (repo *BookRepository) DeleteBook(bookID int) error {
	if _,exists := repo.books[bookID]; !exists {
		return errors.New("id buku tidak ada, jadi tidak bisa di hapus")
	}
	delete(repo.books, bookID)
	return nil
}

// GetAllBooks implements BookRepositoryInterface.
func (repo *BookRepository) GetAllBooks() ([]domain.Buku, error) {
	list := make([]domain.Buku, 0)
	for _,book := range repo.books {
		list = append(list, book)
	}
	return list, nil
}

// GetBookById implements BookRepositoryInterface.
func (repo *BookRepository) GetBookById(bookID int) (*domain.Buku, error) {
	book, exists := repo.books[bookID]
	if !exists {
		return nil, errors.New("id buku tidak di temukan")
	}
	return &book, nil
}

// SaveBook implements BookRepositoryInterface.
func (repo *BookRepository) SaveBook(book *domain.Buku) error {
	if _, exists := repo.books[book.ID]; exists {
		return errors.New("buku sudah ada")
	}

	repo.books[book.ID] = *book
	fmt.Println(repo.books)
	return nil
}

// UpdateBook implements BookRepositoryInterface.
func (repo *BookRepository) UpdateBook(book *domain.Buku) error {
	if _, exists := repo.books[book.ID]; exists {
		repo.books[book.ID] = *book
		return nil
	}
	return nil
}


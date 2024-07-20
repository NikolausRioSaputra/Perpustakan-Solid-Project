package handler

import (
	"Project_Niko/internal/domain"
	"Project_Niko/internal/usecase"
	"errors"
)

type BookHendlerInterface interface {
	StoreNewBook(book domain.Book) error
	UpdateBook(book domain.Book) error
	DeleteBook(book domain.Book) error
	ListBooks() []domain.Book
}

type BookHandler struct {
	BookUsecase usecase.BookUsecaseInterface
}

func NewBookHandler(bookUsecase usecase.BookUsecaseInterface) BookHendlerInterface {
	return BookHandler{
		BookUsecase: bookUsecase,
	}
}

func (h BookHandler) StoreNewBook(book domain.Book) error {
	err := h.BookUsecase.AddBook(book)
	if err != nil {
		return err
	}

	return validateBook(book)
}

func (h BookHandler) UpdateBook(book domain.Book) error {
	err := h.BookUsecase.UpdateBook(book)
	if err != nil {
		return err
	}

	return validateBook(book)
}

func (h BookHandler) DeleteBook(book domain.Book) error {
	err := h.BookUsecase.DeleteBook(book)
	if err != nil {
		return err
	}

	return validateBook(book)
}

func (h BookHandler) ListBooks() []domain.Book {
	err := h.BookUsecase.ListBooks()
	if err != nil {
		return err
	}
	return h.ListBooks()
}

func validateBook(book domain.Book) error {
	if book.ID <= 0 {
		return errors.New("invalid book ID: must be a positive integer")
	}

	if book.Title == "" {
		return errors.New("invalid book title: must be a non-empty string")
	}

	return nil
}
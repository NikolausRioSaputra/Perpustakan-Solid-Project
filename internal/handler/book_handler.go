package handler

import (
	"Project_Niko/internal/domain"
	"Project_Niko/internal/usecase"
)

type BookHendlerInterface interface {
	StoreNewBook(book domain.Book) error
	UpdateBook(book domain.Book) error
	DeleteBook(book domain.Book) error
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

	return nil
}

func (h BookHandler) UpdateBook(book domain.Book) error {
	err := h.BookUsecase.UpdateBook(book)
	if err != nil {
		return err
	}

	return nil
}

func (h BookHandler) DeleteBook(book domain.Book) error {
	err := h.BookUsecase.DeleteBook(book)
	if err != nil {
		return err
	}

	return nil
}

package usecase

import (
	"Project_Niko/internal/domain"
	"Project_Niko/internal/repository"
)

type BookUsecaseInterface interface {
	AddBook
	UpdateBook
	DeleteBook
}

type AddBook interface {
	AddBook(book domain.Book) error
}

type UpdateBook interface {
	UpdateBook(book domain.Book) error
}

type DeleteBook interface {
	DeleteBook(book domain.Book) error
}

type BookUsecaseImpl struct {
	bookRepo repository.BookRepositoryInterface
}

func NewBookUsecase(bookRepo repository.BookRepositoryInterface) BookUsecaseInterface {
	return BookUsecaseImpl{
		bookRepo: bookRepo,
	}
}

func (uc BookUsecaseImpl) AddBook(book domain.Book) error {
	err := uc.bookRepo.SaveBook(&book)
	if err != nil {
		return err
	}

	return nil
}

func (uc BookUsecaseImpl) UpdateBook(book domain.Book) error {
	err := uc.bookRepo.UpdateBook(&book)
	if err != nil {
		return err
	}

	return nil
}

func (uc BookUsecaseImpl) DeleteBook(book domain.Book) error {
	err := uc.bookRepo.DeleteBook(book.ID)
	if err != nil {
		return err
	}

	return nil
}

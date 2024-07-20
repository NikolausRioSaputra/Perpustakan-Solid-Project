package handler

import (
	"Project_Niko/internal/domain"
	"Project_Niko/internal/usecase"
)

type PersonHandlerInterface interface {
	StoreNewPerson(person domain.Person)error
	UpdatePerson(person domain.Person) error
	DeletePerson(person domain.Person) error
}

type PersonHandler struct {
	PersonUsecase usecase.PersonUsecaseInterface
}


func NewPersonHandler(personUsecase usecase.PersonUsecaseInterface) PersonHandlerInterface {
	return PersonHandler{
		PersonUsecase: personUsecase,
	}
}

// DeletePerson implements PersonHandlerInterface.
func (h PersonHandler) DeletePerson(person domain.Person) error {
	err := h.PersonUsecase.DeletePerson(person)
	if err != nil {
		return err
	}
	return nil
}

// StoreNewPerson implements PersonHandlerInterface.
func (h PersonHandler) StoreNewPerson(person domain.Person) error {
	err := h.PersonUsecase.AddBook(person)
	if err != nil {
		return err
	}
	return nil
}

// UpdatePerson implements PersonHandlerInterface.
func (h PersonHandler) UpdatePerson(person domain.Person) error {
	err := h.PersonUsecase.UpdatePerson(person)
	if err != nil {
		return err
	}
	return nil
}


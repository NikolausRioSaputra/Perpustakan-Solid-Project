package usecase

import (
	"Project_Niko/internal/domain"
	"Project_Niko/internal/repository"
)

type PersonUsecaseInterface interface {
	AddPerson
	UpdatePerson
	DeletePerson
	ListPersons
}

type AddPerson interface {
	AddPerson(person domain.Person) error
}

type UpdatePerson interface {
	UpdatePerson(person domain.Person) error
}

type DeletePerson interface {
	DeletePerson(person domain.Person) error
}

type ListPersons interface {
	ListPersons() []domain.Person
}

type PersonUsecase struct {
	personRepo repository.PersonRepositoryInterface
}

func NewPersonUsecase(personRepo repository.PersonRepositoryInterface) PersonUsecaseInterface {
	return PersonUsecase{
		personRepo: personRepo,
	}
}

// AddBook implements PersonUsecaseInterface.
func (uc PersonUsecase) AddPerson(person domain.Person) error {
	err := uc.personRepo.SavaPerson(&person)
	if err != nil {
		return err
	}
	return nil
}

// DeletePerson implements PersonUsecaseInterface.
func (uc PersonUsecase) DeletePerson(person domain.Person) error {
	err := uc.personRepo.DeletePerson(person.ID)
	if err != nil {
		return err
	}
	return nil
}

// UpdatePerson implements PersonUsecaseInterface.
func (uc PersonUsecase) UpdatePerson(person domain.Person) error {
	err := uc.personRepo.UpdatePerson(&person)
	if err != nil {
		return err
	}
	return nil
}

func (uc PersonUsecase) ListPersons() []domain.Person {
	return uc.personRepo.ListPersons()
}

package repository

import (
	"Project_Niko/internal/domain"
	"errors"
)

type PersonRepositoryInterface interface {
	PersonSaver
	PersonUpdate
	PersonDelete
	PersonLister
	PersonChecker
}

type PersonSaver interface {
	SavaPerson(person *domain.Person) error
}

type PersonUpdate interface {
	UpdatePerson(person *domain.Person) error
}

type PersonDelete interface {
	DeletePerson(personID int) error
}

type PersonLister interface {
	ListPersons() []domain.Person
}

type PersonChecker interface {
	IsPersonExists(personID int) bool
}

type PersonRepository struct {
	persons map[int]domain.Person
}

func NewPersonRepository() PersonRepositoryInterface {
	return &PersonRepository{
		persons: map[int]domain.Person{},
	}
}

// DeletePerson implements PersonRepositoryInterface.
func (repo *PersonRepository) DeletePerson(personID int) error {
	if _, exists := repo.persons[personID]; !exists {
		return errors.New("id person not found")
	}
	delete(repo.persons, personID)
	return nil

}

// SavaPerson implements PersonRepositoryInterface.
func (repo *PersonRepository) SavaPerson(person *domain.Person) error {
	if _, exists := repo.persons[person.ID]; exists {
		return errors.New("id person not found")
	}

	repo.persons[person.ID] = *person
	return nil
}

// UpdatePerson implements PersonRepositoryInterface.
func (repo *PersonRepository) UpdatePerson(person *domain.Person) error {
	if _, exists := repo.persons[person.ID]; exists {
		repo.persons[person.ID] = *person
		return nil
	}

	return errors.New("id person not found")
}

func (repo *PersonRepository) ListPersons() []domain.Person {
	persons := []domain.Person{}
	for _, person := range repo.persons {
		persons = append(persons, person)
	}
	return persons
}

func (repo *PersonRepository) IsPersonExists(personID int) bool {
	_, exists := repo.persons[personID]
	return exists
}

package usecase

import (
	"Project_Niko/internal/domain"
	"Project_Niko/internal/repository"
	"errors"
)

type LoanUsecaseInterface interface {
	AddLoan
	UpdateLoan
	DeleteLoan
	ListLoans
	CheckLoan
}

type AddLoan interface {
	AddLoan(loan domain.Loan) error
}

type UpdateLoan interface {
	UpdateLoan(loan domain.Loan) error
}

type DeleteLoan interface {
	DeleteLoan(loan domain.Loan) error
}

type ListLoans interface {
	ListLoans() []domain.Loan
}

type CheckLoan interface {
    IsBookLoaned(bookID int) bool
    IsPersonLoaning(personID int) bool
}

type LoanUsecase struct {
	loanRepo repository.LoanRepositoryInterface
	bookRepo   repository.BookRepositoryInterface
	personRepo repository.PersonRepositoryInterface
}

func NewLoanUsecase(loanRepo repository.LoanRepositoryInterface, bookRepo repository.BookRepositoryInterface, personRepo repository.PersonRepositoryInterface) LoanUsecaseInterface {
	return &LoanUsecase{
		loanRepo: loanRepo,
		bookRepo:   bookRepo,
		personRepo: personRepo,
	}
}

// AddLoan implements LoanUsecaseInterface.
func (uc *LoanUsecase) AddLoan(loan domain.Loan) error {
	if !uc.bookRepo.IsBookExists(loan.Book.ID) {
		return errors.New("book does not exist")
	}

	if !uc.personRepo.IsPersonExists(loan.Person.ID) {
		return errors.New("person does not exist")
	}
	
	err := uc.loanRepo.SaveLoan(&loan)
	if err != nil {
		return err
	}
	return nil
}

// DeleteLoan implements LoanUsecaseInterface.
func (uc *LoanUsecase) DeleteLoan(loan domain.Loan) error {
	err := uc.loanRepo.DeleteLoan(loan.ID)
	if err != nil {
		return err
	}
	return nil
}

// UpdateLoan implements LoanUsecaseInterface.
func (uc *LoanUsecase) UpdateLoan(loan domain.Loan) error {
	err := uc.loanRepo.UpdateLoan(&loan)
	if err != nil {
		return err
	}
	return nil
}

func (uc LoanUsecase) ListLoans() []domain.Loan {
	return uc.loanRepo.ListLoans()
}

func (uc *LoanUsecase) IsBookLoaned(bookID int) bool {
    return uc.loanRepo.IsBookLoaned(bookID)
}

func (uc *LoanUsecase) IsPersonLoaning(personID int) bool {
    return uc.loanRepo.IsPersonLoaning(personID)
}

package repository

import (
	"Project_Niko/internal/domain"
	"errors"
)

type LoanRepositoryInterface interface {
	LoanSaver
	LoanUpdater
	LoanDeleter
	LoanLister
	LoanChecker
}

type LoanSaver interface {
	SaveLoan(loan *domain.Loan) error 
}

type LoanUpdater interface {
	UpdateLoan(loan *domain.Loan) error
}

type LoanDeleter interface {
	DeleteLoan(loadID int) error
}

type LoanLister interface {
	ListLoans() []domain.Loan
}

type LoanChecker interface {
    IsBookLoaned(bookID int) bool
    IsPersonLoaning(personID int) bool
}

type LoanRepository struct {
	loans map[int]domain.Loan
}


func NewLoanRepository() LoanRepositoryInterface {
	return &LoanRepository{
		loans: map[int]domain.Loan{},
	}
}

func (repo *LoanRepository) SaveLoan(loan *domain.Loan) error {
	if _, exists := repo.loans[loan.ID]; exists {
		return errors.New("loan already exist")
	}
	
	repo.loans[loan.ID] = *loan
	return nil
}

func (repo *LoanRepository) UpdateLoan(loan *domain.Loan) error {
	if _, exists := repo.loans[loan.ID]; exists{
		if loan.Returned {
            delete(repo.loans, loan.ID)
        } else {
            repo.loans[loan.ID] = *loan
        }
		return nil
	}

	return errors.New("id loan not found")
}

func (repo *LoanRepository) DeleteLoan(loanID int) error {
	if _,exists := repo.loans[loanID]; !exists {
		return errors.New("id loan not found")
	}

	delete(repo.loans, loanID)
	return nil
}

func (repo *LoanRepository) ListLoans() []domain.Loan {
	loans := []domain.Loan{}
	for _, loan := range repo.loans {
		loans = append(loans, loan)
	}
	return loans
}

func (repo *LoanRepository) IsBookLoaned(bookID int) bool {
    for _, loan := range repo.loans {
        if loan.Book.ID == bookID && !loan.Returned {
            return true
        }
    }
    return false
}

func (repo *LoanRepository) IsPersonLoaning(personID int) bool {
    for _, loan := range repo.loans {
        if loan.Person.ID == personID && !loan.Returned {
            return true
        }
    }
    return false
}
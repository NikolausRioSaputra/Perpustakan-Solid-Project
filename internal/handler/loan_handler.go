package handler

import (
	"Project_Niko/internal/domain"
	"Project_Niko/internal/usecase"
)

type LaonHandlerInterface interface {
	StoreNewLoan(loan domain.Loan) error
	UpdateLoan(loan domain.Loan) error
	DeleteLoan(loan domain.Loan) error
	ListLoans() []domain.Loan
	IsBookLoaned(bookID int) bool
	IsPersonLoaning(personID int) bool
}

type LoanHandler struct {
	LoanUsecase usecase.LoanUsecaseInterface
}

func NewLoanHandler(loanUsecase usecase.LoanUsecaseInterface) LaonHandlerInterface {
	return LoanHandler{
		LoanUsecase: loanUsecase,
	}
}

// DeleteLoan implements LaonHandlerInterface.
func (h LoanHandler) DeleteLoan(loan domain.Loan) error {
	err := h.LoanUsecase.DeleteLoan(loan)
	if err != nil {
		return err
	}
	return nil
}

// StoreNewLoan implements LaonHandlerInterface.
func (h LoanHandler) StoreNewLoan(loan domain.Loan) error {
	err := h.LoanUsecase.AddLoan(loan)
	if err != nil {
		return err
	}
	return nil
}

// UpdateLoan implements LaonHandlerInterface.
func (h LoanHandler) UpdateLoan(loan domain.Loan) error {
	err := h.LoanUsecase.UpdateLoan(loan)
	if err != nil {
		return err
	}
	return nil
}

func (h LoanHandler) ListLoans() []domain.Loan {
	return h.LoanUsecase.ListLoans()
}

func (h LoanHandler) IsBookLoaned(bookID int) bool {
    return h.LoanUsecase.IsBookLoaned(bookID)
}

func (h LoanHandler) IsPersonLoaning(personID int) bool {
    return h.LoanUsecase.IsPersonLoaning(personID)
}
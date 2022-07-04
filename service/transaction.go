package service

import (
	"github.com/Favoree-Team/server-user-api/entity"
	"github.com/Favoree-Team/server-user-api/repository"
)

type TransactionService interface {
	// admin
	GetAllTransaction(page string, limit string) (entity.TransactionPage, error)
	UpdateStatusById(id string, input entity.TransactionStatusInput) error

	// user
	GetTransactionDetail(id string) (entity.Transaction, error)
	CreateTransaction(input entity.RequestTransaction) (entity.Transaction, error)
	ConfirmPaid(id string) error
	CancelTransaction(id string) error
}

type transactionService struct {
	transRepo repository.TransactionRepository
}

func NewTransactionRepository(transRepo repository.TransactionRepository) *transactionService {
	return &transactionService{transRepo: transRepo}
}

func (s *transactionService) GetAllTransaction(page string, limit string) (entity.TransactionPage, error) {
	return entity.TransactionPage{}, nil
}

func (s *transactionService) UpdateStatusById(id string, input entity.TransactionStatusInput) error {
	return nil
}

func (s *transactionService) GetTransactionDetail(id string) (entity.Transaction, error) {
	return entity.Transaction{}, nil
}

func (s *transactionService) CreateTransaction(input entity.RequestTransaction) (entity.Transaction, error) {
	return entity.Transaction{}, nil
}

func (s *transactionService) ConfirmPaid(id string) error {
	return nil
}

func (s *transactionService) CancelTransaction(id string) error {
	return nil
}

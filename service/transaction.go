package service

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/Favoree-Team/server-user-api/config"
	"github.com/Favoree-Team/server-user-api/entity"
	"github.com/Favoree-Team/server-user-api/repository"
	"github.com/Favoree-Team/server-user-api/utils"
)

type TransactionService interface {
	// admin
	GetAllTransaction(page string, limit string) (entity.TransactionPage, error)
	UpdateStatusById(id string, input entity.TransactionStatusInput) error

	// user
	GetTransactionDetail(id string) (entity.Transaction, error)
	CreateTransaction(userId string, input entity.RequestTransaction) (entity.Transaction, error)
	GetLastTransaction(userId string) (entity.LastTransactionResponse, error)
	ConfirmPaid(userId string, id string) error
	CancelTransaction(userId string, id string) error
}

type transactionService struct {
	transRepo repository.TransactionRepository
}

func NewTransactionRepository(transRepo repository.TransactionRepository) *transactionService {
	return &transactionService{transRepo: transRepo}
}

func (s *transactionService) GetAllTransaction(page string, limit string) (entity.TransactionPage, error) {
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		return entity.TransactionPage{}, utils.CreateErrorMsg(http.StatusInternalServerError, err)
	}

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		return entity.TransactionPage{}, utils.CreateErrorMsg(http.StatusInternalServerError, err)
	}

	offset := utils.GetOffsite(pageInt, limitInt)

	transactions, total, err := s.transRepo.GetAll(offset, limitInt)
	if err != nil {
		return entity.TransactionPage{}, utils.CreateErrorMsg(http.StatusInternalServerError, err)
	}

	var transctionPage = entity.TransactionPage{
		Total:       total,
		TotalPage:   total / int64(limitInt),
		CurrentPage: int64(pageInt),
		Data:        transactions.ToListTransactionItemPage(),
	}

	return transctionPage, nil
}

func (s *transactionService) UpdateStatusById(id string, input entity.TransactionStatusInput) error {
	// for admin without check userid

	var edit = map[string]interface{}{
		"updated_at": time.Now(),
		"status":     input.Status,
	}

	if input.Status == string(entity.StatusPending) {
		edit["done"] = false
	} else {
		edit["done"] = true
	}

	err := s.transRepo.UpdateByID(id, edit)
	if err != nil {
		return utils.CreateErrorMsg(http.StatusInternalServerError, err)
	}

	return nil
}

func (s *transactionService) GetTransactionDetail(id string) (entity.Transaction, error) {
	return entity.Transaction{}, nil
}

func (s *transactionService) CreateTransaction(userId string, input entity.RequestTransaction) (entity.Transaction, error) {
	// check last transaction, if user hijack to server

	lastTransaction, err := s.GetLastTransaction(userId)
	if err != nil {
		return entity.Transaction{}, utils.CreateErrorMsg(http.StatusInternalServerError, err)
	}

	if lastTransaction.InternalCode == entity.ErrorInternalCode {
		return entity.Transaction{}, utils.CreateErrorMsg(http.StatusBadRequest, errors.New("last transaction still open, must be paid first"))
	}

	generateId := utils.NewUUID()

	// expired using env
	expiredTime, err := config.GetExpiredTime()
	if err != nil {
		return entity.Transaction{}, utils.CreateErrorMsg(http.StatusInternalServerError, err)
	}

	expired := time.Minute * time.Duration(expiredTime)

	var transaction = entity.Transaction{
		ID:             generateId,
		UserID:         userId,
		SenderNumber:   input.SenderNumber,
		SenderWallet:   input.SenderWallet,
		ReceiverName:   input.ReceiverName,
		ReceiverNumber: input.ReceiverNumber,
		ReceiverWallet: input.ReceiverWallet,
		AmountTransfer: input.AmountTransfer,
		AdminFee:       1000,
		AmountReceived: input.AmountTransfer - 1000,
		Status:         entity.StatusPending,
		Done:           false,
		IsConfirmPaid:  entity.NotConfirmPaid,
		ExpiredAt:      time.Now().Add(expired).String(),
	}

	err = s.transRepo.Insert(transaction)
	if err != nil {
		return entity.Transaction{}, utils.CreateErrorMsg(http.StatusInternalServerError, err)
	}

	return transaction, nil
}

func (s *transactionService) GetLastTransaction(userId string) (entity.LastTransactionResponse, error) {
	var lastTransaction = entity.LastTransactionResponse{
		UserID: userId,
	}

	transactions, err := s.transRepo.GetByUserID(userId)
	if err != nil {
		return entity.LastTransactionResponse{}, utils.CreateErrorMsg(http.StatusInternalServerError, err)
	}

	if len(transactions) == 0 {
		lastTransaction.InternalCode = entity.CreatableInternalCode
		return lastTransaction, nil
	} else {
		lastTransaction.Transaction = transactions[0]

		// jika belum done => error
		// jika belum done tapi sudah confirm paid => bisa create lagi

		if !transactions[0].Done {
			if !transactions[0].IsConfirmPaid {
				lastTransaction.InternalCode = entity.ErrorInternalCode
			} else {
				lastTransaction.InternalCode = entity.CreatableInternalCode
			}

			return lastTransaction, nil
		} else {
			lastTransaction.InternalCode = entity.CreatableInternalCode
			return lastTransaction, nil
		}
	}
}

func (s *transactionService) ConfirmPaid(userId string, id string) error {
	transaction, err := s.transRepo.GetByID(id)
	if err != nil {
		return utils.CreateErrorMsg(http.StatusInternalServerError, err)
	}

	if transaction.UserID != userId {
		return utils.CreateErrorMsg(http.StatusBadRequest, fmt.Errorf("user id %s not have permission to edit", userId))
	}

	var edit = map[string]interface{}{
		"updated_at": time.Now(),
	}

	if transaction.IsConfirmPaid == entity.NotConfirmPaid {
		edit["is_confirm_paid"] = entity.ConfirmPaid

		err := s.transRepo.UpdateByID(id, edit)
		if err != nil {
			return utils.CreateErrorMsg(http.StatusInternalServerError, err)
		} else {
			return nil
		}
	}

	return nil
}

func (s *transactionService) CancelTransaction(userId string, id string) error {
	transaction, err := s.transRepo.GetByID(id)
	if err != nil {
		return utils.CreateErrorMsg(http.StatusInternalServerError, err)
	}

	if transaction.UserID != userId {
		return utils.CreateErrorMsg(http.StatusBadRequest, fmt.Errorf("user id %s not have permission to edit", userId))
	}

	var edit = map[string]interface{}{
		"updated_at": time.Now(),
	}

	edit["status"] = entity.StatusCanceled
	edit["done"] = true

	err = s.transRepo.UpdateByID(id, edit)
	if err != nil {
		return utils.CreateErrorMsg(http.StatusInternalServerError, err)
	} else {
		return nil
	}
}

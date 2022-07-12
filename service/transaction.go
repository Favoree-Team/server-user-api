package service

import (
	"errors"
	"fmt"
	"log"
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
	// GetTransactionDetail(id string) (entity.Transaction, error)
	CreateTransaction(userId string, input entity.RequestTransaction) (entity.Transaction, error)
	GetLastTransaction(userId string) (entity.LastTransactionResponse, error)
	ConfirmPaid(userId string, role string, id string) error
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

	var transactionPage entity.TransactionPage

	if total == 0 {
		transactionPage.TotalPage = 1
		transactionPage.Total = 0
		transactionPage.CurrentPage = int64(pageInt)
		transactionPage.Data = []entity.TransactionItemPage{}
	} else {
		if res := total % int64(limitInt); res == 0 {
			transactionPage.TotalPage = total / int64(limitInt)
		} else {
			transactionPage.TotalPage = total/int64(limitInt) + 1
		}

		transactionPage.Total = total
		transactionPage.CurrentPage = int64(pageInt)
		transactionPage.Data = transactions.ToListTransactionItemPage((pageInt - 1) * limitInt)
	}

	return transactionPage, nil
}

func (s *transactionService) UpdateStatusById(id string, input entity.TransactionStatusInput) error {
	// for admin without check userid

	var edit = map[string]interface{}{
		"updated_at": time.Now(),
		"status":     input.Status,
	}

	if input.Status == entity.StatusPending {
		edit["done"] = false
	} else {
		edit["done"] = true
	}

	if input.Note != "" || len(input.Note) > 0 {
		edit["note"] = entity.GetNoteBody(input.Status) + input.Note
	}

	err := s.transRepo.UpdateByID(id, edit)
	if err != nil {
		return utils.CreateErrorMsg(http.StatusInternalServerError, err)
	}

	return nil
}

// func (s *transactionService) GetTransactionDetail(id string) (entity.Transaction, error) {
// 	return entity.Transaction{}, nil
// }

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

	newestTransaction, err := s.transRepo.GetLastTransactionToday()
	if err != nil {
		return entity.Transaction{}, utils.CreateErrorMsg(http.StatusInternalServerError, err)
	}

	orderID, err := utils.GetOrderNow(newestTransaction)
	if err != nil {
		return entity.Transaction{}, utils.CreateErrorMsg(http.StatusInternalServerError, err)
	}

	var transaction = entity.Transaction{
		ID:             generateId,
		UserID:         userId,
		OrderID:        orderID,
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
		IsConfirmPaid:  false,
		CreatedAt:      time.Now().String(),
		ExpiredAt:      time.Now().Add(expired).String(),
	}

	err = s.transRepo.Insert(transaction)
	if err != nil {
		return entity.Transaction{}, utils.CreateErrorMsg(http.StatusInternalServerError, err)
	}

	//TODO: send notification to email admin

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

		// jika belum expired
		// jika belum done => error
		// jika belum done tapi sudah confirm paid => bisa create lagi

		timeExp, err := utils.ParseStrtoTime(transactions[0].ExpiredAt)
		if err != nil {
			return entity.LastTransactionResponse{}, utils.CreateErrorMsg(http.StatusInternalServerError, err)
		}

		if timeExp.Before(time.Now()) {
			log.Println("masuk case sudah epired")
			lastTransaction.InternalCode = entity.CreatableInternalCode
		} else if !transactions[0].Done && !transactions[0].IsConfirmPaid && transactions[0].Status == entity.StatusPending {
			log.Println("masuk case belum paid, belum done, dan masih pending")
			lastTransaction.InternalCode = entity.ErrorInternalCode
		} else if !transactions[0].Done && transactions[0].IsConfirmPaid {
			log.Println("masuk case belum paid, dan sudah confirm paid")
			lastTransaction.InternalCode = entity.CreatableInternalCode
		} else if transactions[0].Done {
			log.Println("masuk case done")
			lastTransaction.InternalCode = entity.CreatableInternalCode
		}
	}

	return lastTransaction, nil
}

func (s *transactionService) ConfirmPaid(userId string, role string, id string) error {
	transaction, err := s.transRepo.GetByID(id)
	if err != nil {
		return utils.CreateErrorMsg(http.StatusInternalServerError, err)
	}

	if role != "admin" {
		if transaction.UserID != userId {
			return utils.CreateErrorMsg(http.StatusBadRequest, fmt.Errorf("user id %s not have permission to edit", userId))
		}
	}

	var edit = map[string]interface{}{
		"updated_at": time.Now(),
	}

	if !transaction.IsConfirmPaid {
		edit["is_confirm_paid"] = true

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

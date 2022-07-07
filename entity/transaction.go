package entity

import "github.com/go-playground/validator/v10"

type TransactionStatusEnum interface {
	IsValid() bool
}

type TransactionStatus string

type ConfirmPaidStatus bool

const (
	StatusPending TransactionStatus = "pending"

	// user success paid the transaction, after click the upload data in web, update to paid
	StatusPaid TransactionStatus = "paid"

	// success transfer manual from admin, change to success
	// set Done to true
	StatusSuccess TransactionStatus = "success"

	// failed for case user not paid, admin cannot transfer
	// set Done to true
	StatusFailed TransactionStatus = "failed"

	// canceled by user
	// set Done to true
	StatusCanceled TransactionStatus = "canceled"

	// confirmation paid status
	NotConfirmPaid ConfirmPaidStatus = false
	ConfirmPaid    ConfirmPaidStatus = true
)

func (t TransactionStatus) IsValid() bool {
	switch t {
	case StatusPending, StatusPaid, StatusSuccess, StatusFailed, StatusCanceled:
		return true
	}
	return false
}

func ValidateTransStatusEnum(fl validator.FieldLevel) bool {
	return fl.Field().Interface().(TransactionStatusEnum).IsValid()
}

type Transaction struct {
	ID             string            `json:"id"`
	UserID         string            `json:"user_id"`
	SenderNumber   string            `json:"sender_number"`
	SenderWallet   string            `json:"sender_wallet"`
	ReceiverName   string            `json:"receiver_name"`
	ReceiverNumber string            `json:"receiver_number"`
	ReceiverWallet string            `json:"receiver_wallet"`
	AmountTransfer int               `json:"amount_transfer"`
	AdminFee       int               `json:"admin_fee"`
	AmountReceived int               `json:"amount_received"`
	Status         TransactionStatus `json:"status"` // [pending, paid, success, failed, canceled]
	Done           bool              `json:"done"`   // default is false
	IsConfirmPaid  ConfirmPaidStatus `json:"is_confirm_paid"`
	CreatedAt      string            `json:"created_at"`
	UpdatedAt      string            `json:"updated_at"`
	ExpiredAt      string            `json:"expired_at"`
}

type RequestTransaction struct {
	SenderNumber   string `json:"sender_number" binding:"required"`
	SenderWallet   string `json:"sender_wallet" binding:"required"`
	ReceiverName   string `json:"receiver_name" binding:"required"`
	ReceiverNumber string `json:"receiver_number" binding:"required"`
	ReceiverWallet string `json:"receiver_wallet" binding:"required"`
	AmountTransfer int    `json:"amount_transfer" binding:"required"`
}

type TransactionItemPage struct {
	ExternalID     int               `json:"external_id"`
	ID             string            `json:"id"`
	UserID         string            `json:"user_id"`
	SenderNumber   string            `json:"sender_number"`
	SenderWallet   string            `json:"sender_wallet"`
	ReceiverName   string            `json:"receiver_name"`
	ReceiverNumber string            `json:"receiver_number"`
	ReceiverWallet string            `json:"receiver_wallet"`
	AmountTransfer int               `json:"amount_transfer"`
	AdminFee       int               `json:"admin_fee"`
	AmountReceived int               `json:"amount_received"`
	Status         TransactionStatus `json:"status"` // [pending, paid, success, failed, canceled]
	Done           bool              `json:"done"`   // default is false
	IsConfirmPaid  ConfirmPaidStatus `json:"is_confirm_paid"`
	CreatedAt      string            `json:"created_at"`
	UpdatedAt      string            `json:"updated_at"`
	ExpiredAt      string            `json:"expired_at"`
}

// for pagination
type TransactionPage struct {
	Total       int                   `json:"total"`
	TotalPage   int                   `json:"total_page"`
	CurrentPage int                   `json:"current_page"`
	Data        []TransactionItemPage `json:"data"`
}

// edit status by admin
type TransactionStatusInput struct {
	Status string `json:"status" binding:"required,status_enum"` // [pending, paid, success, failed, cancelled]
}

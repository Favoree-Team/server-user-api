package entity

type TransactionStatusEnum interface {
	IsValid() bool
}

type TransactionStatus string

const (
	StatusPending TransactionStatus = "pending"

	// user success paid the transaction, after click the upload data in web, update to paid
	//StatusPaid TransactionStatus = "paid"

	// success transfer manual from admin, change to success
	// set Done to true
	StatusSuccess TransactionStatus = "success"

	// failed for case user not paid, admin cannot transfer
	// set Done to true
	StatusFailed TransactionStatus = "failed"

	// canceled by user
	// set Done to true
	StatusCanceled TransactionStatus = "canceled"
)

var (
	noteBody = map[TransactionStatus]string{
		"pending":  "PENDING : ",
		"success":  "SUCCESS : ",
		"failed":   "FAILED : ",
		"canceled": "CANCELED : ",
	}
)

func (t TransactionStatus) IsValid() bool {
	switch t {
	case StatusPending, StatusSuccess, StatusFailed, StatusCanceled:
		return true
	}
	return false
}

func GetNoteBody(status TransactionStatus) string {
	if _, ok := noteBody[status]; ok {
		return noteBody[status]
	}

	return ""
}

type Transaction struct {
	ID             string            `json:"id"`
	UserID         string            `json:"user_id"`
	OrderID        string            `json:"order_id"`
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
	Note           string            `json:"note"`
	IsConfirmPaid  bool              `json:"is_confirm_paid"`
	CreatedAt      string            `json:"created_at"`
	UpdatedAt      string            `json:"updated_at"`
	ExpiredAt      string            `json:"expired_at"`
}

type ListTransaction []Transaction

func (t ListTransaction) ToListTransactionItemPage(start int) []TransactionItemPage {
	var listTransaction []TransactionItemPage

	if len(t) < 1 {
		return []TransactionItemPage{}
	}

	for i := 0; i < len(t); i++ {
		transactionItemPage := TransactionItemPage{
			ExternalID:     start + (i + 1),
			ID:             t[i].ID,
			UserID:         t[i].UserID,
			OrderID:        t[i].OrderID,
			SenderNumber:   t[i].SenderNumber,
			SenderWallet:   t[i].SenderWallet,
			ReceiverName:   t[i].ReceiverName,
			ReceiverNumber: t[i].ReceiverNumber,
			ReceiverWallet: t[i].ReceiverWallet,
			AmountTransfer: t[i].AmountTransfer,
			AdminFee:       t[i].AdminFee,
			AmountReceived: t[i].AmountReceived,
			Status:         t[i].Status,
			Done:           t[i].Done,
			Note:           t[i].Note,
			IsConfirmPaid:  t[i].IsConfirmPaid,
			CreatedAt:      t[i].CreatedAt,
			UpdatedAt:      t[i].UpdatedAt,
			ExpiredAt:      t[i].ExpiredAt,
		}

		listTransaction = append(listTransaction, transactionItemPage)
	}

	return listTransaction
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
	OrderID        string            `json:"order_id"`
	SenderNumber   string            `json:"sender_number"`
	SenderWallet   string            `json:"sender_wallet"`
	ReceiverName   string            `json:"receiver_name"`
	ReceiverNumber string            `json:"receiver_number"`
	ReceiverWallet string            `json:"receiver_wallet"`
	AmountTransfer int               `json:"amount_transfer"`
	AdminFee       int               `json:"admin_fee"`
	AmountReceived int               `json:"amount_received"`
	Status         TransactionStatus `json:"status"` // [pending, success, failed, canceled]
	Note           string            `json:"note"`
	Done           bool              `json:"done"` // default is false
	IsConfirmPaid  bool              `json:"is_confirm_paid"`
	CreatedAt      string            `json:"created_at"`
	UpdatedAt      string            `json:"updated_at"`
	ExpiredAt      string            `json:"expired_at"`
}

// for pagination
type TransactionPage struct {
	Total       int64                 `json:"total"`
	TotalPage   int64                 `json:"total_page"`
	CurrentPage int64                 `json:"current_page"`
	Data        []TransactionItemPage `json:"data"`
}

// edit status by admin
type TransactionStatusInput struct {
	Status TransactionStatus `json:"status" binding:"required"` // [pending, success, failed, canceled]
	Note   string            `json:"note"`
}

// entity last transaction

type InternalCode string

const (
	ErrorInternalCode     InternalCode = "ERROR"
	CreatableInternalCode InternalCode = "CREATABLE"
)

type LastTransactionResponse struct {
	UserID       string       `json:"user_id"`
	InternalCode InternalCode `json:"internal_code"`
	Transaction  Transaction  `json:"transaction"`
}

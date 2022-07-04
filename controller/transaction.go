package controller

import (
	"github.com/Favoree-Team/server-user-api/service"
	"github.com/gin-gonic/gin"
)

type transactionController struct {
	transService service.TransactionService
}

func NewTransactionController(transService service.TransactionService) *transactionController {
	return &transactionController{
		transService: transService,
	}
}

func (tc *transactionController) GetAllTransaction(c *gin.Context) {
}

func (tc *transactionController) UpdateTransactionByAdmin(c *gin.Context) {
}

func (tc *transactionController) CreateTransaction(c *gin.Context) {

}

func (tc *transactionController) GetTransactionDetail(c *gin.Context) {

}

func (tc *transactionController) ConfirmPaid(c *gin.Context) {

}

func (tc *transactionController) CancelTransaction(c *gin.Context) {

}

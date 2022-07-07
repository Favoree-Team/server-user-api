package controller

import (
	"errors"
	"net/http"

	"github.com/Favoree-Team/server-user-api/entity"
	"github.com/Favoree-Team/server-user-api/service"
	"github.com/Favoree-Team/server-user-api/utils"
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

// =========================== ADMIN ================================

func (tc *transactionController) GetAllTransaction(c *gin.Context) {
	_, ok := c.Get("user_id")
	if !ok {
		errMsg := utils.CreateErrorMsg(http.StatusUnauthorized, errors.New("unauthorize user"))
		c.JSON(http.StatusUnauthorized, utils.ErrorHandler(errMsg))
		return
	}

	role, ok := c.Get("role")
	if !ok {
		errMsg := utils.CreateErrorMsg(http.StatusUnauthorized, errors.New("unauthorize user"))
		c.JSON(http.StatusUnauthorized, utils.ErrorHandler(errMsg))
		return
	}

	if role.(string) != "admin" {
		errMsg := utils.CreateErrorMsg(http.StatusUnauthorized, errors.New("only admin can get all transaction"))
		c.JSON(http.StatusUnauthorized, utils.ErrorHandler(errMsg))
		return
	}

	page := c.Param("page")
	limit := c.Param("limit")

	if page == "" || limit == "" {
		errMsg := utils.CreateErrorMsg(http.StatusBadRequest, errors.New("page and limit is required"))
		c.JSON(http.StatusBadRequest, utils.ErrorHandler(errMsg))
		return
	}

	transactionPage, err := tc.transService.GetAllTransaction(page, limit)
	if err != nil {
		c.JSON(utils.GetErrorCode(err), utils.ErrorHandler(err))
		return
	}

	c.JSON(http.StatusOK, transactionPage)
}

func (tc *transactionController) UpdateTransactionByAdmin(c *gin.Context) {
	_, ok := c.Get("user_id")
	if !ok {
		errMsg := utils.CreateErrorMsg(http.StatusUnauthorized, errors.New("unauthorize user"))
		c.JSON(http.StatusUnauthorized, utils.ErrorHandler(errMsg))
		return
	}

	role, ok := c.Get("role")
	if !ok {
		errMsg := utils.CreateErrorMsg(http.StatusUnauthorized, errors.New("unauthorize user"))
		c.JSON(http.StatusUnauthorized, utils.ErrorHandler(errMsg))
		return
	}

	if role.(string) != "admin" {
		errMsg := utils.CreateErrorMsg(http.StatusUnauthorized, errors.New("only admin can get all transaction"))
		c.JSON(http.StatusUnauthorized, utils.ErrorHandler(errMsg))
		return
	}

	transactionId := c.Param("transaction_id")
	if transactionId == "" {
		errMsg := utils.CreateErrorMsg(http.StatusBadRequest, errors.New("transaction_id is required"))
		c.JSON(http.StatusBadRequest, utils.ErrorHandler(errMsg))
		return
	}

	var input entity.TransactionStatusInput
	if err := c.ShouldBindJSON(&input); err != nil {
		errMsg := utils.CreateErrorMsg(http.StatusBadRequest, err)
		c.JSON(http.StatusBadRequest, utils.ErrorHandler(errMsg))
		return
	}

	err := tc.transService.UpdateStatusById(transactionId, input)
	if err != nil {
		c.JSON(utils.GetErrorCode(err), utils.ErrorHandler(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"transaction_id": transactionId,
		"status":         input.Status,
		"message":        "update transaction status success",
	})
}

// =========================== USER ================================

func (tc *transactionController) CreateTransaction(c *gin.Context) {
	userId, ok := c.Get("user_id")
	if !ok {
		errMsg := utils.CreateErrorMsg(http.StatusUnauthorized, errors.New("unauthorized user"))
		c.JSON(http.StatusUnauthorized, utils.ErrorHandler(errMsg))
		return
	}

	var input entity.RequestTransaction
	if err := c.ShouldBindJSON(&input); err != nil {
		errMsg := utils.CreateErrorMsg(http.StatusBadRequest, err)
		c.JSON(http.StatusBadRequest, utils.ErrorHandler(errMsg))
		return
	}

	transaction, err := tc.transService.CreateTransaction(userId.(string), input)
	if err != nil {
		c.JSON(utils.GetErrorCode(err), utils.ErrorHandler(err))
		return
	}

	c.JSON(http.StatusCreated, transaction)

}

func (tc *transactionController) GetLastTransaction(c *gin.Context) {
	userID, ok := c.Get("user_id")
	if !ok {
		errMsg := utils.CreateErrorMsg(http.StatusUnauthorized, errors.New("unauthorized user"))
		c.JSON(http.StatusUnauthorized, utils.ErrorHandler(errMsg))
		return
	}

	transaction, err := tc.transService.GetLastTransaction(userID.(string))
	if err != nil {
		c.JSON(utils.GetErrorCode(err), utils.ErrorHandler(err))
		return
	}

	c.JSON(http.StatusOK, transaction)
}

// func (tc *transactionController) GetTransactionDetail(c *gin.Context) {

// }

func (tc *transactionController) ConfirmPaid(c *gin.Context) {
	userId, ok := c.Get("user_id")
	if !ok {
		errMsg := utils.CreateErrorMsg(http.StatusUnauthorized, errors.New("unauthorized user"))
		c.JSON(http.StatusUnauthorized, utils.ErrorHandler(errMsg))
		return
	}

	transactionId := c.Param("transaction_id")
	if transactionId == "" {
		errMsg := utils.CreateErrorMsg(http.StatusBadRequest, errors.New("transaction_id is required"))
		c.JSON(http.StatusBadRequest, utils.ErrorHandler(errMsg))
		return
	}

	err := tc.transService.ConfirmPaid(userId.(string), transactionId)
	if err != nil {
		c.JSON(utils.GetErrorCode(err), utils.ErrorHandler(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"transaction_id": transactionId,
		"message":        "success confirm paid",
	})

}

func (tc *transactionController) CancelTransaction(c *gin.Context) {
	userId, ok := c.Get("user_id")
	if !ok {
		errMsg := utils.CreateErrorMsg(http.StatusUnauthorized, errors.New("unauthorized user"))
		c.JSON(http.StatusUnauthorized, utils.ErrorHandler(errMsg))
		return
	}

	transactionId := c.Param("transaction_id")
	if transactionId == "" {
		errMsg := utils.CreateErrorMsg(http.StatusBadRequest, errors.New("transaction_id is required"))
		c.JSON(http.StatusBadRequest, utils.ErrorHandler(errMsg))
		return
	}

	err := tc.transService.CancelTransaction(userId.(string), transactionId)
	if err != nil {
		c.JSON(utils.GetErrorCode(err), utils.ErrorHandler(err))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"transaction_id": transactionId,
		"message":        "success cancel transaction",
	})
}

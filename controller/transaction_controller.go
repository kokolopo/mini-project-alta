package controller

import (
	"errors"
	"net/http"
	"order_kafe/helper"
	"order_kafe/transaction"
	"order_kafe/user"

	"github.com/gin-gonic/gin"
)

type transactionController struct {
	transactionService transaction.TransactionService
	// orderService  order.OrderService
	// detailService detail.DetailOrderService
}

func NewTransactionHandler(transactionService transaction.TransactionService) *transactionController {
	return &transactionController{transactionService}
}

func (ctrl *transactionController) GetUserTransactions(c *gin.Context) {
	// didapatkan dari JWT
	currentUser := c.MustGet("currentUser").(user.User)
	userId := currentUser.ID

	transactions, err := ctrl.transactionService.GetTransactionByUserId(userId)
	if err != nil {
		res := helper.ApiResponse("Data Transactions Not Found", http.StatusBadRequest, "failed", err)

		c.JSON(http.StatusBadRequest, res)
		return
	}

	formatter := transaction.FormatUserTransactions(transactions)
	res := helper.ApiResponse("Fetch All Data of Transactions", http.StatusCreated, "success", formatter)

	c.JSON(http.StatusCreated, res)
}

func (ctrl *transactionController) GetTransactions(c *gin.Context) {
	// cek apakah yg akses adalah admin
	currentUser := c.MustGet("currentUser").(user.User)
	userRole := currentUser.Role
	if userRole != "admin" {
		res := helper.ApiResponse("Failed to Create Menu", http.StatusBadRequest, "failed", errors.New("kamu bukan admin"))

		c.JSON(http.StatusBadRequest, res)
		return
	}

	transactions, err := ctrl.transactionService.GetTransactions()
	if err != nil {
		res := helper.ApiResponse("Data Transactions Not Found", http.StatusBadRequest, "failed", err)

		c.JSON(http.StatusBadRequest, res)
		return
	}

	formatter := transaction.FormatUserTransactions(transactions)
	res := helper.ApiResponse("Fetch All Data of Transactions", http.StatusCreated, "success", formatter)

	c.JSON(http.StatusCreated, res)
}

func (ctrl *transactionController) CreateTransaction(c *gin.Context) {
	var input transaction.InputCreateTansaction

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errMessage := gin.H{"errors": errors}

		res := helper.ApiResponse("Failed to Create Transaction", http.StatusUnprocessableEntity, "failed", errMessage)
		c.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	newTrans, errTrans := ctrl.transactionService.CreateTransaction(input)
	if errTrans != nil {
		res := helper.ApiResponse("Failed to Create Transaction", http.StatusUnprocessableEntity, "failed", errTrans)
		c.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	formatter := transaction.FormatTransaction(newTrans)

	res := helper.ApiResponse("Success to Create Transactions", http.StatusCreated, "success", formatter)
	c.JSON(http.StatusCreated, res)
}

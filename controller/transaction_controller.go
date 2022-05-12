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

package controller

import (
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

	res := helper.ApiResponse("Fetch All Data of Transactions", http.StatusCreated, "success", transactions)

	c.JSON(http.StatusCreated, res)
}

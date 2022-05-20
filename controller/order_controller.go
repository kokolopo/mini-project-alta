package controller

import (
	"errors"
	"net/http"
	detail "order_kafe/detail_order"
	"order_kafe/helper"
	"order_kafe/order"
	"order_kafe/user"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	orderService  order.OrderService
	detailService detail.DetailOrderService
}

func NewOrderHandler(orderService order.OrderService, detailService detail.DetailOrderService) *OrderController {
	return &OrderController{orderService, detailService}
}

func (ctrl *OrderController) CreateNewOrder(c *gin.Context) {
	// didapatkan dari JWT
	currentUser := c.MustGet("currentUser").(user.User)
	userId := currentUser.ID

	var input []detail.InputNewDetailOrder
	// tangkap input body
	err := c.ShouldBindJSON(&input)
	if err != nil {
		res := helper.ApiResponse("New Data Has Been Failed", http.StatusUnprocessableEntity, "failed", err)

		c.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	// record data order
	newOrder, errOrder := ctrl.orderService.CreateOrder(userId)
	if errOrder != nil {
		res := helper.ApiResponse("Order Has Been Failed", http.StatusBadRequest, "failed", errOrder)

		c.JSON(http.StatusBadRequest, res)
	}
	orderId := newOrder.ID

	// record data detail order
	_, errDetals := ctrl.detailService.SaveDetailOrder(orderId, input)
	if errDetals != nil {
		res := helper.ApiResponse("New Data Has Been Failed", http.StatusBadRequest, "failed", errDetals)

		c.JSON(http.StatusBadRequest, res)
	}

	data := gin.H{"is_recorded": true}
	res := helper.ApiResponse("Order Has Been Created", http.StatusCreated, "success", data)

	c.JSON(http.StatusCreated, res)
}

func (ctrl *OrderController) GetUserOrders(c *gin.Context) {
	// cek apakah yg akses adalah admin
	currentUser := c.MustGet("currentUser").(user.User)
	userRole := currentUser.Role
	if userRole != "admin" {
		res := helper.ApiResponse("Failed to Access", http.StatusBadRequest, "failed", errors.New("kamu bukan admin"))

		c.JSON(http.StatusBadRequest, res)
		return
	}

	items, err := ctrl.orderService.GetOrders()
	if err != nil {
		res := helper.ApiResponse("Data Not Found or Error", http.StatusBadRequest, "failed", err)

		c.JSON(http.StatusBadRequest, res)
		return
	}

	formatter := order.FormatUserOrders(items)
	res := helper.ApiResponse("Fetch All Data of User Order", http.StatusOK, "success", formatter)

	c.JSON(http.StatusCreated, res)
}

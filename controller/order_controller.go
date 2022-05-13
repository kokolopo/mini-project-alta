package controller

import (
	"net/http"
	detail "order_kafe/detail_order"
	"order_kafe/helper"
	"order_kafe/order"
	"order_kafe/user"

	"github.com/gin-gonic/gin"
)

type orderController struct {
	orderService  order.OrderService
	detailService detail.DetailOrderService
}

func NewOrderHandler(orderService order.OrderService, detailService detail.DetailOrderService) *orderController {
	return &orderController{orderService, detailService}
}

func (ctrl *orderController) CreateNewOrder(c *gin.Context) {
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
	details, errDetals := ctrl.detailService.SaveDetailOrder(orderId, input)
	if errDetals != nil {
		res := helper.ApiResponse("New Data Has Been Failed", http.StatusBadRequest, "failed", errDetals)

		c.JSON(http.StatusBadRequest, res)
	}

	//formatter := user.Formatitem(newOrder)

	res := helper.ApiResponse("Order Has Been Created", http.StatusCreated, "success", details)

	c.JSON(http.StatusCreated, res)
}

func (ctrl *orderController) GetUserOrders(c *gin.Context) {
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

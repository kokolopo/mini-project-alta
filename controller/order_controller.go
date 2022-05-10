package controller

import (
	"net/http"
	"order_kafe/helper"
	"order_kafe/order"
	"order_kafe/user"

	"github.com/gin-gonic/gin"
)

type orderController struct {
	orderService order.OrderService
}

func NewOrderHandler(orderService order.OrderService) *orderController {
	return &orderController{orderService}
}

func (ctrl *orderController) CreateNewOrder(c *gin.Context) {
	// didapatkan dari JWT
	currentUser := c.MustGet("currentUser").(user.User)
	userId := currentUser.ID

	newOrder, errCate := ctrl.orderService.CreateOrder(userId)
	if errCate != nil {
		res := helper.ApiResponse("Order Has Been Failed", http.StatusBadRequest, "failed", errCate)

		c.JSON(http.StatusBadRequest, res)
	}

	//formatter := user.Formatitem(newOrder)

	res := helper.ApiResponse("Order Has Been Created", http.StatusCreated, "success", newOrder)

	c.JSON(http.StatusCreated, res)
}

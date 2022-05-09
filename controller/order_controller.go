package controller

import (
	"order_kafe/item"
	"order_kafe/order"

	"github.com/gin-gonic/gin"
)

type orderController struct {
	orderService order.OrderService
}

func NewOrderHandler(itemService item.ItemService) *itemController {
	return &itemController{itemService}
}

func (ctrl *orderController) CreateNewOrder(c *gin.Context) {
	// dapatkan id user dari session yang sedang login
	// mapping data body ke struct order detail

}

package controller

import (
	"net/http"
	"order_kafe/helper"
	"order_kafe/item"

	"github.com/gin-gonic/gin"
)

type itemController struct {
	itemService item.ItemService
}

func NewItemHandler(itemService item.ItemService) *itemController {
	return &itemController{itemService}
}

func (ctrl *itemController) CreateNewitem(c *gin.Context) {
	var input item.InputNewItem

	err := c.ShouldBindJSON(&input)
	if err != nil {
		res := helper.ApiResponse("New Data Has Been Failed", http.StatusUnprocessableEntity, "failed", err)

		c.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	newitem, errUser := ctrl.itemService.CreateNewItem(input)
	if errUser != nil {
		res := helper.ApiResponse("New Data Has Been Failed", http.StatusBadRequest, "failed", errUser)

		c.JSON(http.StatusBadRequest, res)
	}

	//formatter := user.Formatitem(newitem)

	res := helper.ApiResponse("New User Data Has Been Created", http.StatusCreated, "success", newitem)

	c.JSON(http.StatusCreated, res)
}

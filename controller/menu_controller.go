package controller

import (
	"net/http"
	"order_kafe/helper"
	"order_kafe/item"
	"strconv"

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

	res := helper.ApiResponse("New Item Has Been Created", http.StatusCreated, "success", newitem)

	c.JSON(http.StatusCreated, res)
}

func (ctrl *itemController) GetItems(c *gin.Context) {
	items, err := ctrl.itemService.GetAllItem()
	if err != nil {
		res := helper.ApiResponse("Data Not Found or Error", http.StatusBadRequest, "failed", err)

		c.JSON(http.StatusBadRequest, res)
		return
	}

	res := helper.ApiResponse("Fetch All Data of Item", http.StatusOK, "success", items)

	c.JSON(http.StatusCreated, res)
}

func (ctrl *itemController) UpdateItems(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var input item.InputUpdateItem
	err := c.ShouldBindJSON(&input)
	if err != nil {
		res := helper.ApiResponse("Update Data Has Been Failed", http.StatusUnprocessableEntity, "failed", err)

		c.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	updatedItem, errUpdate := ctrl.itemService.UpdateItem(id, input)
	if errUpdate != nil {
		res := helper.ApiResponse("Update Data Has Been Failed", http.StatusUnprocessableEntity, "failed", err)

		c.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	res := helper.ApiResponse("Update Data Has Been Success", http.StatusCreated, "success", updatedItem)

	c.JSON(http.StatusCreated, res)
}

func (ctrl *itemController) DeleteItems(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	itemById, err := ctrl.itemService.GetById(id)
	if err != nil {
		res := helper.ApiResponse("Item Not Found", http.StatusBadRequest, "failed", err)

		c.JSON(http.StatusBadRequest, res)
		return
	}

	if itemById.ID == 0 {
		res := helper.ApiResponse("User Not Found", http.StatusBadRequest, "failed", err)

		c.JSON(http.StatusBadRequest, res)
		return
	}

	itemDelete, errDel := ctrl.itemService.DeleteItem(itemById.ID)
	if errDel != nil {
		res := helper.ApiResponse("User Not Found", http.StatusBadRequest, "failed", errDel)

		c.JSON(http.StatusBadRequest, res)
		return
	}

	cekItem, errCek := ctrl.itemService.GetById(id)
	if errCek != nil {
		res := helper.ApiResponse("Any Error", http.StatusBadRequest, "failed", errCek)

		c.JSON(http.StatusBadRequest, res)
		return
	}

	if cekItem.ID == 0 {
		res := helper.ApiResponse("Successfuly Delete User", http.StatusOK, "success", nil)

		c.JSON(http.StatusOK, res)
		return
	}

	res := helper.ApiResponse("Any Error", http.StatusBadRequest, "failed", itemDelete)

	c.JSON(http.StatusCreated, res)
}

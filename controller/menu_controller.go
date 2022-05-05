package controller

import (
	"net/http"
	"order_kafe/helper"
	"order_kafe/menu"

	"github.com/gin-gonic/gin"
)

type menuController struct {
	menuService menu.MenuService
}

func NewMenuHandler(menuService menu.MenuService) *menuController {
	return &menuController{menuService}
}

func (ctrl *menuController) CreateNewMenu(c *gin.Context) {
	var input menu.InputNewMenu

	err := c.ShouldBindJSON(&input)
	if err != nil {
		res := helper.ApiResponse("New Data Has Been Failed", http.StatusUnprocessableEntity, "failed", err)

		c.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	newMenu, errUser := ctrl.menuService.CreateNewMenu(input)
	if errUser != nil {
		res := helper.ApiResponse("New Data Has Been Failed", http.StatusBadRequest, "failed", errUser)

		c.JSON(http.StatusBadRequest, res)
	}

	//formatter := user.FormatMenu(newMenu)

	res := helper.ApiResponse("New User Data Has Been Created", http.StatusCreated, "success", newMenu)

	c.JSON(http.StatusCreated, res)
}

package controller

import (
	"net/http"
	detail "order_kafe/detail_order"
	"order_kafe/helper"

	"github.com/gin-gonic/gin"
)

type detailController struct {
	detailService detail.DetailOrderService
}

func NewDetailHandler(detailService detail.DetailOrderService) *detailController {
	return &detailController{detailService}
}

func (ctrl *detailController) SaveDetailOrder(c *gin.Context) {
	var input []detail.InputNewDetailOrder

	err := c.ShouldBindJSON(&input)
	if err != nil {
		res := helper.ApiResponse("New Data Has Been Failed", http.StatusUnprocessableEntity, "failed", err)

		c.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	details, errDetals := ctrl.detailService.SaveDetailOrder(input)
	if errDetals != nil {
		res := helper.ApiResponse("New Data Has Been Failed", http.StatusBadRequest, "failed", errDetals)

		c.JSON(http.StatusBadRequest, res)
	}

	//formatter := user.Formatitem(newCategory)

	res := helper.ApiResponse("New Category Has Been Created", http.StatusCreated, "success", details)

	c.JSON(http.StatusCreated, res)
}

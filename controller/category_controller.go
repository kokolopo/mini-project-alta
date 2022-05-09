package controller

import (
	"net/http"
	"order_kafe/category"
	"order_kafe/helper"

	"github.com/gin-gonic/gin"
)

type categoryController struct {
	categoryService category.CategoryService
}

func NewCategoryHandler(categoryService category.CategoryService) *categoryController {
	return &categoryController{categoryService}
}

func (ctrl *categoryController) CreateNewCategory(c *gin.Context) {
	var input category.InputNewCategory

	err := c.ShouldBindJSON(&input)
	if err != nil {
		res := helper.ApiResponse("New Data Has Been Failed", http.StatusUnprocessableEntity, "failed", err)

		c.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	newCategory, errCate := ctrl.categoryService.CreateNewCategory(input)
	if errCate != nil {
		res := helper.ApiResponse("New Data Has Been Failed", http.StatusBadRequest, "failed", errCate)

		c.JSON(http.StatusBadRequest, res)
	}

	//formatter := user.Formatitem(newCategory)

	res := helper.ApiResponse("New Category Has Been Created", http.StatusCreated, "success", newCategory)

	c.JSON(http.StatusCreated, res)
}

func (ctrl *categoryController) GetCategories(c *gin.Context) {
	categories, err := ctrl.categoryService.GetAllICategory()
	if err != nil {
		res := helper.ApiResponse("Data Not Found or Error", http.StatusBadRequest, "failed", err)

		c.JSON(http.StatusBadRequest, res)
		return
	}

	res := helper.ApiResponse("Fetch All Data of Category", http.StatusOK, "success", categories)

	c.JSON(http.StatusCreated, res)
}

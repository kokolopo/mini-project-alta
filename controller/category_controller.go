package controller

import (
	"errors"
	"net/http"
	"order_kafe/category"
	"order_kafe/helper"
	"order_kafe/user"
	"strconv"

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

	// cek apakah yg akses adalah admin
	currentUser := c.MustGet("currentUser").(user.User)
	userRole := currentUser.Role
	if userRole != "admin" {
		res := helper.ApiResponse("Failed to Access", http.StatusBadRequest, "failed", errors.New("kamu bukan admin"))

		c.JSON(http.StatusBadRequest, res)
		return
	}

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

	//formatter :=

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

func (ctrl *categoryController) DeleteCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	// cek apakah yg akses adalah admin
	currentUser := c.MustGet("currentUser").(user.User)
	userRole := currentUser.Role
	if userRole != "admin" {
		res := helper.ApiResponse("Failed to Access", http.StatusBadRequest, "failed", errors.New("kamu bukan admin"))

		c.JSON(http.StatusBadRequest, res)
		return
	}

	itemById, err := ctrl.categoryService.GetCategoryById(id)
	if err != nil {
		res := helper.ApiResponse("Item Not Found", http.StatusBadRequest, "failed", err)

		c.JSON(http.StatusBadRequest, res)
		return
	}

	if itemById.ID == 0 {
		res := helper.ApiResponse("Category Not Found", http.StatusBadRequest, "failed", err)

		c.JSON(http.StatusBadRequest, res)
		return
	}

	itemDelete, errDel := ctrl.categoryService.DeleteCategory(itemById.ID)
	if errDel != nil {
		res := helper.ApiResponse("Category Not Found", http.StatusBadRequest, "failed", errDel)

		c.JSON(http.StatusBadRequest, res)
		return
	}

	cekItem, errCek := ctrl.categoryService.GetCategoryById(id)
	if errCek != nil {
		res := helper.ApiResponse("Any Error", http.StatusBadRequest, "failed", errCek)

		c.JSON(http.StatusBadRequest, res)
		return
	}

	if cekItem.ID == 0 {
		res := helper.ApiResponse("Successfuly Delete Category", http.StatusOK, "success", nil)

		c.JSON(http.StatusOK, res)
		return
	}

	res := helper.ApiResponse("Any Error", http.StatusBadRequest, "failed", itemDelete)

	c.JSON(http.StatusCreated, res)
}

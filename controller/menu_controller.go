package controller

import (
	"errors"
	"fmt"
	"net/http"
	"order_kafe/helper"
	"order_kafe/item"
	"order_kafe/user"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ItemController struct {
	itemService item.ItemService
}

func NewItemHandler(itemService item.ItemService) *ItemController {
	return &ItemController{itemService}
}

func (ctrl *ItemController) CreateNewItem(c *gin.Context) {
	// cek apakah yg akses adalah admin
	currentUser := c.MustGet("currentUser").(user.User)
	userRole := currentUser.Role
	if userRole != "admin" {
		res := helper.ApiResponse("Failed to Create Menu", http.StatusBadRequest, "failed", errors.New("kamu bukan admin"))

		c.JSON(http.StatusBadRequest, res)
		return
	}

	var input item.InputNewItem

	err := c.ShouldBindJSON(&input)
	if err != nil {
		res := helper.ApiResponse("Failed create new menu", http.StatusUnprocessableEntity, "failed", err)

		c.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	_, errItem := ctrl.itemService.CreateNewItem(input)
	if errItem != nil {
		res := helper.ApiResponse("New Data Has Been Failed", http.StatusBadRequest, "failed", errItem)

		c.JSON(http.StatusBadRequest, res)
		return
	}

	data := gin.H{"is_created": true}
	res := helper.ApiResponse("New Item Has Been Created", http.StatusCreated, "success", data)

	c.JSON(http.StatusCreated, res)
}

func (ctrl *ItemController) GetItems(c *gin.Context) {
	items, err := ctrl.itemService.GetAllItem()
	if err != nil {
		res := helper.ApiResponse("Data Not Found or Error", http.StatusBadRequest, "failed", err)

		c.JSON(http.StatusBadRequest, res)
		return
	}

	formatter := item.FormatItems(items)
	res := helper.ApiResponse("Fetch All Data of Item", http.StatusOK, "success", formatter)

	c.JSON(http.StatusCreated, res)
}

func (ctrl *ItemController) UpdateItems(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	// cek apakah yg akses adalah admin
	currentUser := c.MustGet("currentUser").(user.User)
	userRole := currentUser.Role
	if userRole != "admin" {
		res := helper.ApiResponse("Failed to Upload Avatar Image", http.StatusBadRequest, "failed", errors.New("kamu bukan admin"))

		c.JSON(http.StatusBadRequest, res)
		return
	}

	var input item.InputUpdateItem
	err := c.ShouldBindJSON(&input)
	if err != nil {
		res := helper.ApiResponse("Update Data Has Been Failed", http.StatusUnprocessableEntity, "failed", err)

		c.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	_, errUpdate := ctrl.itemService.UpdateItem(id, input)
	if errUpdate != nil {
		res := helper.ApiResponse("Update Data Has Been Failed", http.StatusUnprocessableEntity, "failed", err)

		c.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	data := gin.H{"is_uploaded": true}
	res := helper.ApiResponse("Update Data Has Been Success", http.StatusCreated, "success", data)

	c.JSON(http.StatusCreated, res)
}

func (ctrl *ItemController) DeleteItems(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	// cek apakah yg akses adalah admin
	currentUser := c.MustGet("currentUser").(user.User)
	userRole := currentUser.Role
	if userRole != "admin" {
		res := helper.ApiResponse("Failed to Upload Avatar Image", http.StatusBadRequest, "failed", errors.New("kamu bukan admin"))

		c.JSON(http.StatusBadRequest, res)
		return
	}

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

	_, errDel := ctrl.itemService.DeleteItem(itemById.ID)
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
		res := helper.ApiResponse("Successfuly Delete Item", http.StatusOK, "success", nil)

		c.JSON(http.StatusOK, res)
		return
	}

	data := gin.H{"is_deleted": true}
	res := helper.ApiResponse("Any Error", http.StatusBadRequest, "failed", data)

	c.JSON(http.StatusCreated, res)
}

func (h *ItemController) UploadImage(c *gin.Context) {
	item_id, _ := strconv.Atoi(c.Param("id"))

	// cek apakah yg akses adalah admin
	currentUser := c.MustGet("currentUser").(user.User)
	userRole := currentUser.Role
	if userRole != "admin" {
		res := helper.ApiResponse("Failed to Upload Image", http.StatusBadRequest, "failed", errors.New("kamu bukan admin"))

		c.JSON(http.StatusBadRequest, res)
		return
	}

	// cek apakah item ada
	itemById, err := h.itemService.GetById(item_id)
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

	// get image dari form file
	file, err := c.FormFile("image")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		res := helper.ApiResponse("Failed to Upload Image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, res)
		return
	}

	path := fmt.Sprintf("images/%s", file.Filename)

	errImage := c.SaveUploadedFile(file, path)
	if errImage != nil {
		data := gin.H{"is_uploaded": false}
		res := helper.ApiResponse("Failed to Upload Image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, res)
		return
	}

	_, errItem := h.itemService.SaveImage(item_id, path)
	if errItem != nil {
		data := gin.H{"is_uploaded": false}
		res := helper.ApiResponse("Failed to Upload Image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, res)
		return
	}

	data := gin.H{"is_uploaded": true}
	res := helper.ApiResponse("Avatar Successfuly Uploaded", http.StatusOK, "success", data)

	c.JSON(http.StatusOK, res)
}

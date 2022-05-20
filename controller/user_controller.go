package controller

import (
	"fmt"
	"net/http"
	"order_kafe/auth"
	"order_kafe/helper"
	"order_kafe/user"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService user.UserService
	authService auth.Service
}

func NewUserHandler(userService user.UserService, authService auth.Service) *UserController {
	return &UserController{userService, authService}
}

func (ctrl *UserController) UserRegister(c *gin.Context) {
	var input user.InputRegister
	err := c.ShouldBindJSON(&input)
	if err != nil {
		res := helper.ApiResponse("New User Data Has Been Failed", http.StatusUnprocessableEntity, "failed", err)

		c.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	newUser, errUser := ctrl.userService.Register(input)
	if errUser != nil {
		res := helper.ApiResponse("New User Data Has Been Failed", http.StatusBadRequest, "failed", errUser)

		c.JSON(http.StatusBadRequest, res)
		return
	}

	token, errToken := ctrl.authService.GenerateTokenJWT(newUser.ID, newUser.Fullname, newUser.Role)

	if errToken != nil {
		res := helper.ApiResponse("New User Data Has Been Failed", http.StatusBadRequest, "failed", errUser)

		c.JSON(http.StatusBadRequest, res)
		return
	}

	formatter := user.FormatUser(newUser, token)

	res := helper.ApiResponse("New User Data Has Been Created", http.StatusCreated, "success", formatter)

	c.JSON(http.StatusCreated, res)
}

func (h *UserController) Login(c *gin.Context) {
	var input user.InputLogin

	err := c.ShouldBindJSON(&input)
	if err != nil {
		res := helper.ApiResponse("Login Failed", http.StatusUnprocessableEntity, "failed", err)
		c.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	loginUser, errLogin := h.userService.Login(input)
	if errLogin != nil {
		res := helper.ApiResponse("Login Gagal", http.StatusUnprocessableEntity, "failed", errLogin)

		c.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	token, errToken := h.authService.GenerateTokenJWT(loginUser.ID, loginUser.Fullname, loginUser.Role)

	if errToken != nil {
		res := helper.ApiResponse("Fail Create Token", http.StatusBadRequest, "failed", errToken)

		c.JSON(http.StatusBadRequest, res)
		return
	}

	formatter := user.FormatUser(loginUser, token)

	res := helper.ApiResponse("berhasil login", http.StatusOK, "success", formatter)

	c.JSON(http.StatusCreated, res)
}

func (h *UserController) CheckEmailAvailability(c *gin.Context) {
	var input user.InputCheckEmail

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.ApiResponse("Email checking failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	isEmailAvailable, err := h.userService.IsEmailAvailable(input)
	if err != nil {
		errorMessage := gin.H{"errors": "Server error"}
		response := helper.ApiResponse("Email checking failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data := gin.H{
		"is_available": isEmailAvailable,
	}

	metaMessage := "Email has been registered"

	if isEmailAvailable {
		metaMessage = "Email is available"
	}

	response := helper.ApiResponse(metaMessage, http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}

func (h *UserController) UploadAvatar(c *gin.Context) {
	file, err := c.FormFile("avatar")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		res := helper.ApiResponse("Failed to Upload Avatar Image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, res)
		return
	}

	// didapatkan dari JWT
	currentUser := c.MustGet("currentUser").(user.User)
	userId := currentUser.ID

	path := fmt.Sprintf("images/%d-%s", userId, file.Filename)

	errImage := c.SaveUploadedFile(file, path)
	if errImage != nil {
		data := gin.H{"is_uploaded": false}
		res := helper.ApiResponse("Failed to Upload Avatar Image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, res)
		return
	}

	_, errUser := h.userService.SaveAvatar(userId, path)
	if errUser != nil {
		data := gin.H{"is_uploaded": false}
		res := helper.ApiResponse("Failed to Upload Avatar Image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, res)
		return
	}

	data := gin.H{"is_uploaded": true}
	res := helper.ApiResponse("Avatar Successfuly Uploaded", http.StatusOK, "success", data)

	c.JSON(http.StatusOK, res)
}

func (h *UserController) UpdateData(c *gin.Context) {
	// cek yg akses login
	currentUser := c.MustGet("currentUser").(user.User)
	userId := currentUser.ID

	var input user.InputUpdate
	err := c.ShouldBindJSON(&input)
	if err != nil {
		res := helper.ApiResponse("Update Data Has Been Failed", http.StatusUnprocessableEntity, "failed", err)

		c.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	updatedItem, errUpdate := h.userService.UpdateUser(userId, input)
	if errUpdate != nil {
		res := helper.ApiResponse("Update Data Has Been Failed", http.StatusUnprocessableEntity, "failed", err)

		c.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	res := helper.ApiResponse("Update Data Has Been Success", http.StatusCreated, "success", updatedItem)

	c.JSON(http.StatusCreated, res)
}

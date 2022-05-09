package controller

import (
	"fmt"
	"net/http"
	"order_kafe/auth"
	"order_kafe/helper"
	"order_kafe/user"
	"strconv"

	"github.com/gin-gonic/gin"
)

type userController struct {
	userService user.UserService
	authService auth.Service
}

func NewUserHandler(userService user.UserService, authService auth.Service) *userController {
	return &userController{userService, authService}
}

func (ctrl *userController) UserRegister(c *gin.Context) {
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

func (h *userController) Login(c *gin.Context) {
	var input user.InputLogin

	err := c.ShouldBindJSON(&input)
	if err != nil {
		res := helper.ApiResponse("Login Failed", http.StatusUnprocessableEntity, "failed", err)
		c.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	loginUser, errLogin := h.userService.Login(input)
	if errLogin != nil {
		res := helper.ApiResponse("Login Failed", http.StatusUnprocessableEntity, "failed", errLogin)

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

	res := helper.ApiResponse("Successfuly Login", http.StatusOK, "success", formatter)

	c.JSON(http.StatusCreated, res)
}

func (h *userController) CheckEmailAvailability(c *gin.Context) {
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

func (h *userController) UploadAvatar(c *gin.Context) {
	file, err := c.FormFile("avatar")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		res := helper.ApiResponse("Failed to Upload Avatar Image", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, res)
		return
	}

	// harusnya didapatkan dari JWT
	userId := 1

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
	res := helper.ApiResponse("Avatar Successfuly Uploaded", http.StatusBadRequest, "success", data)

	c.JSON(http.StatusOK, res)
}

func (h *userController) UpdateData(c *gin.Context) {
	idParam := c.Param("id")
	id, _ := strconv.Atoi(idParam)

	var input user.InputUpdate

	err := c.ShouldBind(&input)
	if err != nil {
		input.Error = err
		c.HTML(http.StatusOK, "user_edit.html", input)
		return
	}

	input.ID = id

	_, err = h.userService.UpdateUser(input)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error.html", nil)
		return
	}

	c.Redirect(http.StatusFound, "/users")
}

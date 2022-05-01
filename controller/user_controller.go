package controller

import (
	"net/http"
	"order_kafe/auth"
	"order_kafe/helper"
	"order_kafe/user"

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
	err := c.Bind(&input)
	if err != nil {
		res := helper.ApiResponse("New User Data Has Been Failed", http.StatusUnprocessableEntity, "failed", err)

		c.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	newUser, errUser := ctrl.userService.Register(input)
	if errUser != nil {
		res := helper.ApiResponse("New User Data Has Been Failed", http.StatusBadRequest, "failed", errUser)

		c.JSON(http.StatusBadRequest, res)
	}

	token, errToken := ctrl.authService.GenerateTokenJWT(newUser.ID, newUser.NamaLengkap)
	if errToken != nil {
		res := helper.ApiResponse("New User Data Has Been Failed", http.StatusBadRequest, "failed", errUser)

		c.JSON(http.StatusBadRequest, res)
	}

	formatter := user.FormatUser(newUser, token)

	res := helper.ApiResponse("New User Data Has Been Created", http.StatusCreated, "success", formatter)

	c.JSON(http.StatusCreated, res)
}

func (h *userController) Login(c *gin.Context) {
	var input user.InputLogin

	err := c.Bind(&input)
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

	token, errToken := h.authService.GenerateTokenJWT(loginUser.ID, loginUser.NamaLengkap)
	if errToken != nil {
		res := helper.ApiResponse("New User Data Has Been Failed", http.StatusBadRequest, "failed", errToken)

		c.JSON(http.StatusBadRequest, res)
		return
	}

	formatter := user.FormatUser(loginUser, token)

	res := helper.ApiResponse("Successfuly Login", http.StatusOK, "success", formatter)

	c.JSON(http.StatusCreated, res)
}

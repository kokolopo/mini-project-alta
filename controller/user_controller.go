package controller

import "order_kafe/user"

type userController struct {
	userService user.UserService
}

func NewUserHandler(userService user.UserService) *userController {
	return &userController{userService}
}

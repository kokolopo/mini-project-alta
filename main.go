package main

import (
	"order_kafe/auth"
	"order_kafe/config"
	"order_kafe/controller"
	"order_kafe/database"
	"order_kafe/user"

	"github.com/gin-gonic/gin"
)

func main() {
	conf := config.InitConfiguration()
	database.InitDatabase(conf)
	db := database.DB

	userRepo := user.NewUserRepository(db)
	userService := user.NewUserService(userRepo)
	authService := auth.NewService()
	userController := controller.NewUserHandler(userService, authService)

	router := gin.Default()
	api := router.Group("/api/v1")
	api.POST("/register", userController.UserRegister)
	api.POST("/login", userController.Login)

	router.Run(":8080")
}

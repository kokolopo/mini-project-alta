package main

import (
	"order_kafe/auth"
	"order_kafe/config"
	"order_kafe/controller"
	"order_kafe/database"
	"order_kafe/item"
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

	itemRepo := item.NewItemRepository(db)
	itemService := item.NewItemService(itemRepo)
	itemController := controller.NewItemHandler(itemService)

	router := gin.Default()
	api := router.Group("/api/v1")

	// User domain
	api.POST("/users", userController.UserRegister)
	api.POST("/sessions", userController.Login)
	api.POST("/email_checkers", userController.CheckEmailAvailability)
	api.PUT("/users/:id", userController.UpdateData)

	// item domain
	api.POST("/item", itemController.CreateNewitem)
	router.Run(":8080")
}

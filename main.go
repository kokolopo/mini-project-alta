package main

import (
	"order_kafe/auth"
	"order_kafe/category"
	"order_kafe/config"
	"order_kafe/controller"
	"order_kafe/database"
	detail "order_kafe/detail_order"
	"order_kafe/item"
	"order_kafe/order"
	"order_kafe/payment"
	"order_kafe/transaction"
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

	categoryRepo := category.NewCategoryRepository(db)
	categoryService := category.NewCategoryService(categoryRepo)
	categoryController := controller.NewCategoryHandler(categoryService)

	itemRepo := item.NewItemRepository(db)
	itemService := item.NewItemService(itemRepo, categoryRepo)
	itemController := controller.NewItemHandler(itemService)

	orderRepo := order.NewOrderRepository(db)
	detailRepo := detail.NewDetailOrderRepository(db)
	detailService := detail.NewDetailOrderService(detailRepo)
	orderService := order.NewOrderService(orderRepo, userRepo)
	orderController := controller.NewOrderHandler(orderService, detailService)

	transactionRepo := transaction.NewTransactionRepository(db)
	paymentService := payment.NewService()
	transactionService := transaction.NewTransactionService(transactionRepo, paymentService, userRepo)
	transactionController := controller.NewTransactionHandler(transactionService)

	router := gin.Default()
	api := router.Group("/api/v1")

	// user domain
	api.POST("/users", userController.UserRegister)
	api.POST("/sessions", userController.Login)
	api.POST("/email_checkers", userController.CheckEmailAvailability)
	api.PUT("/users", auth.AuthMiddleware(authService, userService), userController.UpdateData)
	api.POST("/avatars", auth.AuthMiddleware(authService, userService), userController.UploadAvatar)

	// item domain
	api.POST("/items", auth.AuthMiddleware(authService, userService), itemController.CreateNewItem)
	api.POST("/items/:id", auth.AuthMiddleware(authService, userService), itemController.UploadImage)
	api.GET("/items", itemController.GetItems)
	api.PUT("/items/:id", auth.AuthMiddleware(authService, userService), itemController.UpdateItems)
	api.DELETE("/items/:id", auth.AuthMiddleware(authService, userService), itemController.DeleteItems)

	// category domain
	api.POST("/categories", auth.AuthMiddleware(authService, userService), categoryController.CreateNewCategory)
	api.GET("/categories", categoryController.GetCategories)
	api.DELETE("/categories/:id", auth.AuthMiddleware(authService, userService), categoryController.DeleteCategory)

	// order domain
	api.POST("/orders", auth.AuthMiddleware(authService, userService), orderController.CreateNewOrder)
	api.GET("/orders", auth.AuthMiddleware(authService, userService), orderController.GetUserOrders)

	// transaction domain
	api.GET("users/transactions", auth.AuthMiddleware(authService, userService), transactionController.GetUserTransactions)
	api.GET("/transactions", auth.AuthMiddleware(authService, userService), transactionController.GetTransactions)
	api.POST("/transactions", auth.AuthMiddleware(authService, userService), transactionController.CreateTransaction)

	router.Run(":8080")
}

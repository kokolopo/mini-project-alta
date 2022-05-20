package routes

import (
	"order_kafe/auth"
	"order_kafe/rest"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ApiRoutes(db *gorm.DB) {
	// auth
	authService := auth.NewService()

	// user domain
	userController, userService := rest.UserDomainLayer(db)

	// category domain
	categoryController := rest.CategoryDomain(db)

	// item domain
	itemController := rest.ItemDomainLayer(db)

	// order domain
	orderController := rest.OrderDomain(db)

	// transaction domain
	transactionController := rest.TransactionDomain(db)

	router := gin.Default()
	api := router.Group("/api/v1")

	// user
	api.POST("/users", userController.UserRegister)
	api.POST("/sessions", userController.Login)
	api.POST("/email_checkers", userController.CheckEmailAvailability)
	api.PUT("/users", auth.AuthMiddleware(authService, userService), userController.UpdateData)
	api.POST("/avatars", auth.AuthMiddleware(authService, userService), userController.UploadAvatar)

	// item
	api.POST("/items", auth.AuthMiddleware(authService, userService), itemController.CreateNewItem)
	api.POST("/items/:id", auth.AuthMiddleware(authService, userService), itemController.UploadImage)
	api.GET("/items", itemController.GetItems)
	api.PUT("/items/:id", auth.AuthMiddleware(authService, userService), itemController.UpdateItems)
	api.DELETE("/items/:id", auth.AuthMiddleware(authService, userService), itemController.DeleteItems)

	// category
	api.POST("/categories", auth.AuthMiddleware(authService, userService), categoryController.CreateNewCategory)
	api.GET("/categories", categoryController.GetCategories)
	api.DELETE("/categories/:id", auth.AuthMiddleware(authService, userService), categoryController.DeleteCategory)

	// order
	api.POST("/orders", auth.AuthMiddleware(authService, userService), orderController.CreateNewOrder)
	api.GET("/orders", auth.AuthMiddleware(authService, userService), orderController.GetUserOrders)

	// transaction
	api.GET("users/transactions", auth.AuthMiddleware(authService, userService), transactionController.GetUserTransactions)
	api.GET("/transactions", auth.AuthMiddleware(authService, userService), transactionController.GetTransactions)
	api.POST("/transactions", auth.AuthMiddleware(authService, userService), transactionController.CreateTransaction)
	api.POST("/transactions/notification", transactionController.GetNotification)

	router.Run(":8080")
}

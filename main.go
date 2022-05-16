package main

import (
	"net/http"
	"order_kafe/auth"
	"order_kafe/category"
	"order_kafe/config"
	"order_kafe/controller"
	"order_kafe/database"
	detail "order_kafe/detail_order"
	"order_kafe/helper"
	"order_kafe/item"
	"order_kafe/order"
	"order_kafe/payment"
	"order_kafe/transaction"
	"order_kafe/user"
	"strings"

	"github.com/dgrijalva/jwt-go"
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
	api.PUT("/users", authMiddleware(authService, userService), userController.UpdateData)
	api.POST("/avatars", authMiddleware(authService, userService), userController.UploadAvatar)

	// item domain
	api.POST("/items", authMiddleware(authService, userService), itemController.CreateNewItem)
	api.POST("/items/:id", authMiddleware(authService, userService), itemController.UploadImage)
	api.GET("/items", itemController.GetItems)
	api.PUT("/items/:id", authMiddleware(authService, userService), itemController.UpdateItems)
	api.DELETE("/items/:id", authMiddleware(authService, userService), itemController.DeleteItems)

	// category domain
	api.POST("/categories", authMiddleware(authService, userService), categoryController.CreateNewCategory)
	api.GET("/categories", categoryController.GetCategories)
	api.DELETE("/categories/:id", authMiddleware(authService, userService), categoryController.DeleteCategory)

	// order domain
	api.POST("/orders", authMiddleware(authService, userService), orderController.CreateNewOrder)
	api.GET("/orders", authMiddleware(authService, userService), orderController.GetUserOrders)

	// transaction domain
	api.GET("user/transactions", authMiddleware(authService, userService), transactionController.GetUserTransactions)
	api.GET("/transactions", authMiddleware(authService, userService), transactionController.GetTransactions)
	api.POST("/transactions", authMiddleware(authService, userService), transactionController.CreateTransaction)

	router.Run(":8080")
}

func authMiddleware(authService auth.Service, userService user.UserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.ApiResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.ApiResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			response := helper.ApiResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userID := int(claim["id"].(float64))

		user, err := userService.GetUserById(userID)
		if err != nil {
			response := helper.ApiResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUser", user)
	}
}

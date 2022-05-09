package main

import (
	"net/http"
	"order_kafe/auth"
	"order_kafe/config"
	"order_kafe/database"
	"order_kafe/helper"
	"order_kafe/order"
	"order_kafe/user"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func main() {
	conf := config.InitConfiguration()
	database.InitDatabase(conf)
	db := database.DB

	orderRepo := order.NewOrderRepository(db)
	userRepo := user.NewUserRepository(db)
	orderService := order.NewOrderService(orderRepo, userRepo)

	var input order.InputNewOrder
	input.UserID = 7
	input.Infomation = "wait a payment"

	orderService.CreateOrder(input)

	// userRepo := user.NewUserRepository(db)
	// userService := user.NewUserService(userRepo)
	// authService := auth.NewService()
	// userController := controller.NewUserHandler(userService, authService)

	// categoryRepo := category.NewCategoryRepository(db)
	// categoryService := category.NewCategoryService(categoryRepo)
	// categoryController := controller.NewCategoryHandler(categoryService)

	// itemRepo := item.NewItemRepository(db)
	// itemService := item.NewItemService(itemRepo, categoryRepo)
	// itemController := controller.NewItemHandler(itemService)

	// router := gin.Default()
	// api := router.Group("/api/v1")

	// // user domain
	// api.POST("/users", userController.UserRegister)
	// api.POST("/sessions", userController.Login)
	// api.POST("/email_checkers", userController.CheckEmailAvailability)
	// api.PUT("/users/:id", userController.UpdateData)
	// api.POST("/avatars", authMiddleware(authService, userService), userController.UploadAvatar)

	// // item domain
	// api.POST("/items", authMiddleware(authService, userService), itemController.CreateNewItem)
	// api.GET("/items", itemController.GetItems)
	// api.PUT("/items/:id", itemController.UpdateItems)
	// api.DELETE("/items/:id", itemController.DeleteItems)

	// // category domain
	// api.POST("/categories", categoryController.CreateNewCategory)
	// api.GET("/categories", categoryController.GetCategories)

	// router.Run(":8080")
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

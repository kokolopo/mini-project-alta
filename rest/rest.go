package rest

import (
	"order_kafe/auth"
	"order_kafe/category"
	"order_kafe/controller"
	detail "order_kafe/detail_order"
	"order_kafe/item"
	"order_kafe/order"
	"order_kafe/payment"
	"order_kafe/transaction"
	"order_kafe/user"

	"gorm.io/gorm"
)

func UserDomainLayer(db *gorm.DB) (*controller.UserController, user.UserService) {
	userRepo := user.NewUserRepository(db)
	userService := user.NewUserService(userRepo)
	authService := auth.NewService()
	userController := controller.NewUserHandler(userService, authService)

	return userController, userService
}

func ItemDomainLayer(db *gorm.DB) *controller.ItemController {
	itemRepo := item.NewItemRepository(db)
	categoryRepo := category.NewCategoryRepository(db)
	itemService := item.NewItemService(itemRepo, categoryRepo)
	itemController := controller.NewItemHandler(itemService)

	return itemController
}

func CategoryDomain(db *gorm.DB) *controller.CategoryController {
	categoryRepo := category.NewCategoryRepository(db)
	categoryService := category.NewCategoryService(categoryRepo)
	categoryController := controller.NewCategoryHandler(categoryService)

	return categoryController
}

func OrderDomain(db *gorm.DB) *controller.OrderController {
	userRepo := user.NewUserRepository(db)
	orderRepo := order.NewOrderRepository(db)
	detailRepo := detail.NewDetailOrderRepository(db)
	detailService := detail.NewDetailOrderService(detailRepo)
	orderService := order.NewOrderService(orderRepo, userRepo)
	orderController := controller.NewOrderHandler(orderService, detailService)

	return orderController
}

func TransactionDomain(db *gorm.DB) *controller.TransactionController {
	transactionRepo := transaction.NewTransactionRepository(db)
	userRepo := user.NewUserRepository(db)
	orderRepo := order.NewOrderRepository(db)
	paymentService := payment.NewService()
	transactionService := transaction.NewTransactionService(transactionRepo, paymentService, userRepo, orderRepo)
	transactionController := controller.NewTransactionHandler(transactionService)

	return transactionController
}

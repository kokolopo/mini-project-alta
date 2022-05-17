package transaction

import "order_kafe/user"

type InputCreateTansaction struct {
	User    user.User
	OrderID int `json:"order_id" binding:"required"`
	Amount  int `json:"amount" binding:"required"`
}

type InputTransactionNotif struct {
	TransactionStatus string `json:"transaction_status"`
	OrderID           string `json:"order_id"`
	PaymentType       string `json:"payment_type"`
	FraudStatus       string `json:"fraud_status"`
}

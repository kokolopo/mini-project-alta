package transaction

import "order_kafe/user"

type InputCreateTansaction struct {
	User    user.User
	OrderID int `json:"order_id" binding:"required"`
	Amount  int `json:"amount" binding:"required"`
}

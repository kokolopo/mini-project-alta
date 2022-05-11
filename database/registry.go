package database

import (
	"order_kafe/category"
	detail "order_kafe/detail_order"
	"order_kafe/item"
	"order_kafe/order"
	"order_kafe/transaction"
	"order_kafe/user"
)

type Model struct {
	Model interface{}
}

func RegisterModel() []Model {
	return []Model{
		{Model: user.User{}},
		{Model: item.Item{}},
		{Model: category.Categorie{}},
		{Model: order.Order{}},
		{Model: detail.DetailOrder{}},
		{Model: transaction.Transaction{}},
	}
}

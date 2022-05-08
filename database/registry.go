package database

import (
	"order_kafe/item"
	"order_kafe/user"
)

type Model struct {
	Model interface{}
}

func RegisterModel() []Model {
	return []Model{
		{Model: user.User{}}, {Model: item.Item{}},
	}
}

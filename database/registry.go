package database

import (
	"order_kafe/menu"
	"order_kafe/user"
)

type Model struct {
	Model interface{}
}

func RegisterModel() []Model {
	return []Model{
		{Model: user.User{}}, {Model: menu.Menu{}},
	}
}

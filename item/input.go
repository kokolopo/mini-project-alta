package item

import "order_kafe/user"

type InputNewItem struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
	Category    string  `json:"category" binding:"required"`
}

type InputUpdateItem struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Category    string  `json:"category"`
	IsAvailable int     `json:"is_available"`
	User        user.User
	//ImageUrl    string  `json:"image_url"`
}

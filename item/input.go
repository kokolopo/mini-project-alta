package item

import (
	"order_kafe/category"
	"order_kafe/user"
)

type InputNewItem struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
	CategoryID  int     `json:"category_id" binding:"required"`
	Category    category.Categorie
}

type InputUpdateItem struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	CategoryID  int     `json:"category_id"`
	IsAvailable int     `json:"is_available"`
	User        user.User
	//ImageUrl    string  `json:"image_url"`
}

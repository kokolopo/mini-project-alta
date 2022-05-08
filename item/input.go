package item

type InputNewItem struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Price       float64 `json:"harga" binding:"required"`
	Category    string  `json:"category" binding:"required"`
}

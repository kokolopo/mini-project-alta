package detail

type InputNewDetailOrder struct {
	OrderID  int    `json:"order_id" binding:"required"`
	ItemID   int    `json:"item_id" binding:"required"`
	Quantity int    `json:"quantity" binding:"required"`
	Note     string `json:"note" binding:"required"`
}

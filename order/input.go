package order

type InputNewOrder struct {
	UserID     int    `json:"user_id" binding:"required"`
	Infomation string `json:"infomation" binding:"required"`
}

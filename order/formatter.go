package order

type UserOrderFormatter struct {
	ID         int              `json:"id"`
	Fullname   string           `json:"fullname"`
	Infomation string           `json:"infomation"`
	Details    []OrderFormatter `json:"details"`
}

type OrderFormatter struct {
	Name     string `json:"item_name"`
	Quantity int    `json:"quantity"`
	Note     string `json:"note"`
}

func FormatUserOrder(order Order) UserOrderFormatter {
	formatter := UserOrderFormatter{}

	formatter.ID = order.ID
	formatter.Fullname = order.User.Fullname
	formatter.Infomation = order.Infomation
	formatter.Details = FormatOrder(order)

	return formatter
}

func FormatUserOrders(order []Order) []UserOrderFormatter {
	if len(order) == 0 {
		return []UserOrderFormatter{}
	}

	var orderFormatter []UserOrderFormatter

	for _, transaction := range order {
		formatter := FormatUserOrder(transaction)
		orderFormatter = append(orderFormatter, formatter)
	}

	return orderFormatter
}

func FormatOrder(order Order) []OrderFormatter {
	var detail []OrderFormatter

	for _, v := range order.Details {
		detail = append(detail, OrderFormatter{Name: v.Item.Name, Quantity: v.Quantity, Note: v.Note})
	}

	return detail
}

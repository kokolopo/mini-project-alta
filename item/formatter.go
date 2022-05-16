package item

type ItemFormatter struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Category    string `json:"category"`
}

func FormatItem(item Item) ItemFormatter {
	formatter := ItemFormatter{}

	formatter.ID = item.ID
	formatter.Name = item.Name
	formatter.Description = item.Description
	formatter.Price = int(item.Price)
	formatter.Category = item.Category.Name

	return formatter
}

func FormatItems(items []Item) []ItemFormatter {
	if len(items) == 0 {
		return []ItemFormatter{}
	}

	var itemsFormatter []ItemFormatter

	for _, item := range items {
		formatter := FormatItem(item)
		itemsFormatter = append(itemsFormatter, formatter)
	}

	return itemsFormatter
}

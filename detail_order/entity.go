package detailorder

import (
	"order_kafe/item"
	"order_kafe/order"
	"time"
)

type DetailOrder struct {
	ID        int         `gorm:"primary_key;auto_increment;not_null"`
	OrderID   int         `gorm:"type:int(25);not null"`
	Order     order.Order `gorm:"foreignKey:OrderID;not null"`
	ItemID    int         `gorm:"type:int(25);not null"`
	Item      item.Item   `gorm:"foreignKey:ItemID;not null"`
	Quantity  int         `gorm:"type:int(25);not null"`
	Note      string      `gorm:"type:longtext"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

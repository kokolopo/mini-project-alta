package order

import (
	"order_kafe/item"
	"order_kafe/user"
	"time"
)

type Order struct {
	ID         int       `gorm:"primary_key;auto_increment;not_null"`
	UserID     int       `gorm:"type:int(25);not null"`
	User       user.User `gorm:"foreignKey:UserID;not null"`
	Infomation string    `gorm:"type:longtext;not null"`
	Details    []DetailOrder
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type DetailOrder struct {
	ID        int       `gorm:"primary_key;auto_increment;not_null"`
	OrderID   int       `gorm:"type:int(25);not null"`
	Order     Order     `gorm:"foreignKey:OrderID;not null"`
	ItemID    int       `gorm:"type:int(25);not null"`
	Item      item.Item `gorm:"foreignKey:ItemID;not null"`
	Quantity  int       `gorm:"type:int(25);not null"`
	Note      string    `gorm:"type:longtext"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

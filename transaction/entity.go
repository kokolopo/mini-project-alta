package transaction

import (
	"order_kafe/order"
	"order_kafe/user"
	"time"
)

type Transaction struct {
	ID        int         `gorm:"primary_key;auto_increment;not_null"`
	UserID    int         `gorm:"type:int(25);not null"`
	User      user.User   `gorm:"foreignKey:UserID;not null"`
	OrderID   int         `gorm:"type:int(25);not null"`
	Order     order.Order `gorm:"foreignKey:OrderID;not null"`
	Amount    int         `gorm:"type:int(100);not null"`
	Status    string      `gorm:"type:varchar(50);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

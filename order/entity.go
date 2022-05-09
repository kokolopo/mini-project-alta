package order

import (
	"order_kafe/user"
	"time"
)

type Order struct {
	ID         int       `gorm:"primary_key;auto_increment;not_null"`
	UserID     int       `gorm:"type:int(25);not null"`
	User       user.User `gorm:"foreignKey:UserID;not null"`
	Infomation string    `gorm:"type:longtext;not null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

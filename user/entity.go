package user

import "time"

type User struct {
	ID        int    `gorm:"primary_key;auto_increment;not_null"`
	Fullname  string `gorm:"type:varchar(50);not null"`
	Email     string `gorm:"type:varchar(100);not null"`
	Password  string `gorm:"type:longtext;not null"`
	Role      string `gorm:"type:varchar(10);not null"`
	Avatar    string `gorm:"type:varchar(255)"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

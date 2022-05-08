package item

import "time"

type Item struct {
	ID          int     `gorm:"primary_key;auto_increment;not_null"`
	Name        string  `gorm:"type:varchar(50);not null"`
	Description string  `gorm:"type:longtext;not null"`
	Price       float64 `gorm:"type:float;not null"`
	Category    string  `gorm:"type:varchar(50);not null"`
	ImageUrl    string  `gorm:"type:varchar(255)"`
	IsAvailable int     `gorm:"type:int(5)"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

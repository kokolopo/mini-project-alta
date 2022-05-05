package menu

import "time"

type Menu struct {
	ID             int     `gorm:"primary_key;auto_increment;not_null"`
	Nama           string  `gorm:"type:varchar(50);not null"`
	Deskripsi      string  `gorm:"type:longtext;not null"`
	Harga          float64 `gorm:"type:float;not null"`
	Kategori       string  `gorm:"type:varchar(50);not null"`
	UrlGambar      string  `gorm:"type:varchar(255)"`
	ApakahTersedia int     `gorm:"type:int(5)"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

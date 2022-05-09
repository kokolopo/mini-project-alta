package category

import "time"

type Categorie struct {
	ID        int    `gorm:"primary_key;auto_increment;not_null"`
	Name      string `gorm:"type:varchar(50);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

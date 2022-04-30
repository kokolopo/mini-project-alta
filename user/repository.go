package user

import "gorm.io/gorm"

type UserRepository interface {
}

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

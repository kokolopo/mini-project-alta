package menu

import "gorm.io/gorm"

type MenuRepository interface {
	Save(menu Menu) (Menu, error)
	FindById(id int) (Menu, error)
	Update(menu Menu) (Menu, error)
	Delete(id int) (Menu, error)
}

type menuRepository struct {
	DB *gorm.DB
}

func NewMenuRepository(db *gorm.DB) *menuRepository {
	return &menuRepository{db}
}

func (r *menuRepository) Save(menu Menu) (Menu, error) {
	err := r.DB.Create(&menu).Error
	if err != nil {
		return menu, err
	}

	return menu, nil
}

func (r *menuRepository) FindById(id int) (Menu, error) {
	var menu Menu
	err := r.DB.Where("id = ?", id).Find(&menu).Error
	if err != nil {
		return menu, err
	}

	return menu, nil
}

func (r *menuRepository) Update(menu Menu) (Menu, error) {
	err := r.DB.Save(&menu).Error

	if err != nil {
		return menu, err
	}

	return menu, nil
}

func (r *menuRepository) Delete(id int) (Menu, error) {
	var menu Menu

	err := r.DB.Where("id = ?", id).Delete(&menu).Error
	if err != nil {
		return menu, err
	}

	return menu, nil
}

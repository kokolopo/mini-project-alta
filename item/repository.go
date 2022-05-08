package item

import "gorm.io/gorm"

type ItemRepository interface {
	Save(item Item) (Item, error)
	FindById(id int) (Item, error)
	Update(item Item) (Item, error)
	Delete(id int) (Item, error)
}

type itemRepository struct {
	DB *gorm.DB
}

func NewItemRepository(db *gorm.DB) *itemRepository {
	return &itemRepository{db}
}

func (r *itemRepository) Save(item Item) (Item, error) {
	err := r.DB.Create(&item).Error
	if err != nil {
		return item, err
	}

	return item, nil
}

func (r *itemRepository) FindById(id int) (Item, error) {
	var item Item
	err := r.DB.Where("id = ?", id).Find(&item).Error
	if err != nil {
		return item, err
	}

	return item, nil
}

func (r *itemRepository) Update(item Item) (Item, error) {
	err := r.DB.Save(&item).Error

	if err != nil {
		return item, err
	}

	return item, nil
}

func (r *itemRepository) Delete(id int) (Item, error) {
	var item Item

	err := r.DB.Where("id = ?", id).Delete(&item).Error
	if err != nil {
		return item, err
	}

	return item, nil
}

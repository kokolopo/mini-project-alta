package item

import "gorm.io/gorm"

type ItemRepository interface {
	Save(item Item) (Item, error)
	FetchAll() ([]Item, error)
	FindById(id int) (Item, error)
	Update(item Item) (Item, error)
	Delete(item Item) (Item, error)
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

func (r *itemRepository) FetchAll() ([]Item, error) {
	var items []Item

	err := r.DB.Preload("Category").Find(&items).Error
	if err != nil {
		return items, err
	}

	return items, nil
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

func (r *itemRepository) Delete(item Item) (Item, error) {
	err := r.DB.Delete(&item).Error
	if err != nil {
		return item, err
	}

	return item, nil
}

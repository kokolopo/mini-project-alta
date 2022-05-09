package item

import "order_kafe/category"

type ItemService interface {
	CreateNewItem(input InputNewItem) (Item, error)
	GetAllItem() ([]Item, error)
	GetById(id int) (Item, error)
	UpdateItem(id int, input InputUpdateItem) (Item, error)
	DeleteItem(id int) (Item, error)
}

type itemService struct {
	repository         ItemRepository
	categoryRepository category.CategoryRepository
}

func NewItemService(repository ItemRepository, categoryRepository category.CategoryRepository) *itemService {
	return &itemService{repository, categoryRepository}
}

func (s *itemService) CreateNewItem(input InputNewItem) (Item, error) {
	var item Item

	//tangkap nilai dari inputan
	item.Name = input.Name
	item.Description = input.Description
	item.Price = input.Price
	item.CategoryID = input.CategoryID
	item.ImageUrl = "Default.jpg"
	item.IsAvailable = 1

	//cek apakah categori ada
	category, errCate := s.categoryRepository.FindById(input.CategoryID)
	if errCate != nil {
		return item, errCate
	}

	item.Category = category

	//save data yang sudah dimapping kedalam struct Mahasiswa
	newItem, err := s.repository.Save(item)
	if err != nil {
		return newItem, err
	}

	return newItem, nil
}

func (s *itemService) GetAllItem() ([]Item, error) {
	items, err := s.repository.FetchAll()
	if err != nil {
		return items, err
	}

	return items, nil
}

func (s *itemService) GetById(id int) (Item, error) {
	item, err := s.repository.FindById(id)
	if err != nil {
		return item, err
	}

	return item, nil

}

func (s *itemService) UpdateItem(id int, input InputUpdateItem) (Item, error) {
	item, errItem := s.repository.FindById(id)
	if errItem != nil {
		return item, errItem
	}

	// if input.User.Role != "admin" {
	// 	return item, errors.New("Not An Admin")
	// }

	item.Name = input.Name
	item.Description = input.Description
	item.Price = input.Price
	item.CategoryID = input.CategoryID
	item.IsAvailable = input.IsAvailable

	updatedItem, errUpdate := s.repository.Update(item)
	if errUpdate != nil {
		return updatedItem, errUpdate
	}

	return updatedItem, nil
}

func (s *itemService) DeleteItem(id int) (Item, error) {
	item, err := s.repository.FindById(id)
	if err != nil {
		return item, err
	}

	deleteItem, errDel := s.repository.Delete(item)
	if errDel != nil {
		return deleteItem, errDel
	}

	return deleteItem, nil
}

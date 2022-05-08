package item

type ItemService interface {
	CreateNewItem(input InputNewItem) (Item, error)
}

type itemService struct {
	repository ItemRepository
}

func NewItemService(repository ItemRepository) *itemService {
	return &itemService{repository}
}

func (s *itemService) CreateNewItem(input InputNewItem) (Item, error) {
	var item Item

	//tangkap nilai dari inputan
	item.Name = input.Name
	item.Description = input.Description
	item.Price = input.Price
	item.Category = input.Category
	item.ImageUrl = "Default.jpg"
	item.IsAvailable = 1

	//save data yang sudah dimapping kedalam struct Mahasiswa
	newItem, err := s.repository.Save(item)
	if err != nil {
		return newItem, err
	}

	return newItem, nil
}

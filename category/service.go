package category

type CategoryService interface {
	CreateNewCategory(input InputNewCategory) (Categorie, error)
	GetAllICategory() ([]Categorie, error)
	// GetById(id int) (Categorie, error)
	// UpdateItem(id int, input InputNewCategory) (Categorie, error)
	// DeleteItem(id int) (Categorie, error)
}

type categoryService struct {
	repository CategoryRepository
}

func NewCategoryService(repository CategoryRepository) *categoryService {
	return &categoryService{repository}
}

func (s *categoryService) CreateNewCategory(input InputNewCategory) (Categorie, error) {
	var category Categorie

	//tangkap nilai dari inputan
	category.Name = input.Name

	//save data yang sudah dimapping kedalam struct Mahasiswa
	newCate, err := s.repository.Save(category)
	if err != nil {
		return newCate, err
	}

	return newCate, nil
}

func (s *categoryService) GetAllICategory() ([]Categorie, error) {
	categories, err := s.repository.FetchAll()
	if err != nil {
		return categories, err
	}

	return categories, nil
}

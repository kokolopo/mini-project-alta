package category

import "gorm.io/gorm"

type CategoryRepository interface {
	Save(category Categorie) (Categorie, error)
	FetchAll() ([]Categorie, error)
	FindById(id int) (Categorie, error)
	Delete(categorie Categorie) (Categorie, error)
}

type categoryRepository struct {
	DB *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) *categoryRepository {
	return &categoryRepository{db}
}

func (r *categoryRepository) Save(category Categorie) (Categorie, error) {
	err := r.DB.Create(&category).Error
	if err != nil {
		return category, err
	}

	return category, nil
}

func (r *categoryRepository) FetchAll() ([]Categorie, error) {
	var categories []Categorie

	err := r.DB.Find(&categories).Error
	if err != nil {
		return categories, err
	}

	return categories, nil
}

func (r *categoryRepository) FindById(id int) (Categorie, error) {
	var category Categorie
	err := r.DB.Where("id = ?", id).Find(&category).Error
	if err != nil {
		return category, err
	}

	return category, nil
}

func (r *categoryRepository) Delete(categorie Categorie) (Categorie, error) {
	err := r.DB.Delete(&categorie).Error
	if err != nil {
		return categorie, err
	}

	return categorie, nil
}

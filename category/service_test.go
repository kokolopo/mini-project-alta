package category

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var repoCategory = &RepositoryMock{Mock: mock.Mock{}}
var serviceCategory = categoryService{repository: repoCategory}

// CreateNewCategory(input InputNewCategory) (Categorie, error)
// GetAllICategory() ([]Categorie, error)
// GetCategoryById(id int) (Categorie, error)
// DeleteCategory(id int) (Categorie, error)

func TestDeleteCategory(t *testing.T) {
	category := Categorie{
		ID:        1,
		Name:      "jus",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	repoCategory.Mock.On("FindById", 1).Return(category)
	repoCategory.Mock.On("Delete", category).Return(nil)

	result, _ := serviceCategory.DeleteCategory(1)
	assert.Nil(t, nil)
	assert.NotNil(t, result)
}

func TestGetCategoryById(t *testing.T) {
	category := Categorie{
		ID:        1,
		Name:      "jus",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	repoCategory.Mock.On("FindById", 2).Return(category)

	result, err := serviceCategory.GetCategoryById(2)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, category.ID, result.ID)
	assert.Equal(t, category.Name, result.Name)
}

func TestGetAllICategory(t *testing.T) {
	var categories = []Categorie{
		{
			ID:        1,
			Name:      "snack",
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
		},
		{
			ID:        2,
			Name:      "jus",
			CreatedAt: time.Time{},
			UpdatedAt: time.Time{},
		},
	}

	repoCategory.Mock.On("FetchAll").Return(categories)

	result, err := serviceCategory.GetAllICategory()
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, categories[0].ID, result[0].ID)
	assert.Equal(t, categories[0].Name, result[0].Name)
}

func TestCreateNewCategory(t *testing.T) {
	var input = InputNewCategory{
		Name: "sea food",
	}
	var category = Categorie{
		ID:   0,
		Name: input.Name,
	}

	repoCategory.Mock.On("Save", category).Return(category)

	result, err := serviceCategory.CreateNewCategory(input)
	assert.Nil(t, err)
	assert.NotNil(t, category)
	assert.Equal(t, category.ID, result.ID)
	assert.Equal(t, category.Name, result.Name)
}

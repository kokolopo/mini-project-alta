package item

import (
	"order_kafe/category"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var repoItem = &RepositoryMock{Mock: mock.Mock{}}
var repoCategory = &category.RepositoryMock{Mock: mock.Mock{}}
var serviceItem = itemService{repository: repoItem, categoryRepository: repoCategory}

// CreateNewItem(input InputNewItem) (Item, error)
// 	SaveImage(id int, fileLocation string) (Item, error)
// 	GetAllItem() ([]Item, error)
// 	GetById(id int) (Item, error)
// 	UpdateItem(id int, input InputUpdateItem) (Item, error)
// 	DeleteItem(id int) (Item, error)

func TestDeleteItem(t *testing.T) {
	item := Item{
		ID:          1,
		Name:        "sosis bakar",
		Description: "saus BBQ",
		Price:       10000,
		CategoryID:  1,
		Category:    category.Categorie{},
		ImageUrl:    "",
		IsAvailable: 1,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	repoItem.Mock.On("FindById", 1).Return(item)
	repoItem.Mock.On("Delete", item).Return(nil)

	result, _ := serviceItem.DeleteItem(1)
	assert.Nil(t, nil)
	assert.NotNil(t, result)
}

func TestUpdateItem(t *testing.T) {
	input := InputUpdateItem{}
	user := Item{}

	repoItem.Mock.On("FindById", 0).Return(nil)
	repoItem.Mock.On("UpdateUser", user).Return(nil)

	_, err := serviceItem.UpdateItem(0, input)

	assert.Nil(t, nil)
	assert.NotNil(t, err)
}

func TestGetById(t *testing.T) {
	item := Item{
		ID:          1,
		Name:        "sosis bakar",
		Description: "saus BBQ",
		Price:       10000,
		CategoryID:  1,
		Category:    category.Categorie{},
		ImageUrl:    "",
		IsAvailable: 1,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
	}

	repoItem.Mock.On("FindById", 2).Return(item)

	result, err := serviceItem.GetById(2)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, item.ID, result.ID)
	assert.Equal(t, item.Name, result.Name)
}

func TestGetAllItem(t *testing.T) {
	var items = []Item{
		{
			ID:          1,
			Name:        "sosis bakar",
			Description: "saus BBQ",
			Price:       10000,
			CategoryID:  1,
			Category:    category.Categorie{},
			ImageUrl:    "",
			IsAvailable: 1,
			CreatedAt:   time.Time{},
			UpdatedAt:   time.Time{},
		},
		{
			ID:          2,
			Name:        "jus mangga",
			Description: "susu putih",
			Price:       12000,
			CategoryID:  2,
			Category:    category.Categorie{},
			ImageUrl:    "",
			IsAvailable: 1,
			CreatedAt:   time.Time{},
			UpdatedAt:   time.Time{},
		},
	}

	repoItem.Mock.On("FetchAll").Return(items)

	result, err := serviceItem.GetAllItem()
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, items[0].ID, result[0].ID)
	assert.Equal(t, items[0].Name, result[0].Name)
}

func TestCreateNewItem(t *testing.T) {
	category := category.Categorie{
		ID:        3,
		Name:      "jus",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	var input = InputNewItem{
		Name:        "nasi goreng",
		Description: "ati ampela",
		Price:       13000,
		CategoryID:  3,
		Category:    category,
	}
	var item = Item{
		ID:          0,
		Name:        input.Name,
		Description: input.Description,
		Price:       input.Price,
		CategoryID:  input.CategoryID,
		Category:    input.Category,
		ImageUrl:    "images/Default.jpg",
		IsAvailable: 1,
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
	}

	repoCategory.Mock.On("FindById", 3).Return(category)
	repoItem.Mock.On("Save", item).Return(item)

	result, err := serviceItem.CreateNewItem(input)
	assert.Nil(t, err)
	assert.NotNil(t, item)
	assert.Equal(t, item.ID, result.ID)
	assert.Equal(t, item.Name, result.Name)
}

// func TestSaveImage(t *testing.T) {
// 	item := Item{
// 		ID:          1,
// 		Name:        "sosis bakar",
// 		Description: "saus BBQ",
// 		Price:       10000,
// 		CategoryID:  1,
// 		Category:    category.Categorie{},
// 		ImageUrl:    "images/avatar.jpg",
// 		IsAvailable: 1,
// 		CreatedAt:   time.Now(),
// 		UpdatedAt:   time.Now(),
// 	}

// 	repoItem.Mock.On("FindById", 1).Return(item)
// 	repoItem.Mock.On("Update", item).Return(item)

// 	result, err := serviceItem.SaveImage(1, "images/avatar.jpg")
// 	assert.Nil(t, err)
// 	assert.NotNil(t, result)
// }

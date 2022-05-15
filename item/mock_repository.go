package item

import (
	"errors"

	"github.com/stretchr/testify/mock"
)

type RepositoryMock struct {
	Mock mock.Mock
}

func NewRepositoryMock(mock mock.Mock) *RepositoryMock {
	return &RepositoryMock{mock}
}

// Save(item Item) (Item, error)
// 	FetchAll() ([]Item, error)
// 	FindById(id int) (Item, error)
// 	Update(item Item) (Item, error)
// 	Delete(item Item) (Item, error)

func (r *RepositoryMock) Save(item Item) (Item, error) {
	argument := r.Mock.Called(item)
	if argument.Get(0) == nil {
		return item, errors.New("ada yang salah")
	} else {
		newItem := argument.Get(0).(Item)
		return newItem, nil
	}
}

func (r *RepositoryMock) FetchAll() ([]Item, error) {

	argument := r.Mock.Called()
	if argument.Get(0) == nil {
		return []Item{}, errors.New("ada yang salah")
	} else {
		item := argument.Get(0).([]Item)
		return item, nil
	}
}

func (r *RepositoryMock) FindById(id int) (Item, error) {

	argument := r.Mock.Called(id)
	if argument.Get(0) == nil {
		return Item{}, errors.New("ada yang salah")
	} else {
		item := argument.Get(0).(Item)
		return item, nil
	}
}

func (r *RepositoryMock) Delete(item Item) (Item, error) {
	argument := r.Mock.Called(item)
	if argument.Get(0) == nil {
		return item, errors.New("ada yang salah")
	} else {
		newUser := argument.Get(0).(Item)
		return newUser, nil
	}
}

func (r *RepositoryMock) Update(item Item) (Item, error) {
	argument := r.Mock.Called(item)
	if argument.Get(0) == nil {
		return item, errors.New("ada yang salah")
	} else {
		item := argument.Get(0).(Item)
		return item, nil
	}
}

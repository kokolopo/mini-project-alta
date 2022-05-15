package category

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

// Save(category Categorie) (Categorie, error)
// FetchAll() ([]Categorie, error)
// FindById(id int) (Categorie, error)
// Delete(categorie Categorie) (Categorie, error)

func (r *RepositoryMock) Save(category Categorie) (Categorie, error) {
	argument := r.Mock.Called(category)
	if argument.Get(0) == nil {
		return category, errors.New("ada yang salah")
	} else {
		newUser := argument.Get(0).(Categorie)
		return newUser, nil
	}
}

func (r *RepositoryMock) FetchAll() ([]Categorie, error) {

	argument := r.Mock.Called()
	if argument.Get(0) == nil {
		return []Categorie{}, errors.New("ada yang salah")
	} else {
		category := argument.Get(0).([]Categorie)
		return category, nil
	}
}

func (r *RepositoryMock) FindById(id int) (Categorie, error) {

	argument := r.Mock.Called(id)
	if argument.Get(0) == nil {
		return Categorie{}, errors.New("ada yang salah")
	} else {
		newUser := argument.Get(0).(Categorie)
		return newUser, nil
	}
}

func (r *RepositoryMock) Delete(categorie Categorie) (Categorie, error) {
	argument := r.Mock.Called(categorie)
	if argument.Get(0) == nil {
		return categorie, errors.New("ada yang salah")
	} else {
		newUser := argument.Get(0).(Categorie)
		return newUser, nil
	}
}

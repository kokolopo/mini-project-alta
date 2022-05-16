package user

import (
	"errors"

	"github.com/stretchr/testify/mock"
)

// type RepositoryMock interface {
// 	Save(user User) (User, error)
// 	FindById(id int) (User, error)
// 	FindByEmail(email string) (User, error)
// 	//UpdateOneByID(id int, user User) (User, error)
// 	Update(user User) (User, error)
// }

type RepositoryMock struct {
	Mock mock.Mock
}

func NewRepositoryMock(mock mock.Mock) *RepositoryMock {
	return &RepositoryMock{mock}
}

func (r *RepositoryMock) Save(user User) (User, error) {

	argument := r.Mock.Called(user)
	if argument.Get(0) == nil {
		return user, errors.New("ada yang salah")
	} else {
		newUser := argument.Get(0).(User)
		return newUser, nil
	}
}

func (r *RepositoryMock) FindById(id int) (User, error) {
	argument := r.Mock.Called(id)
	if argument.Get(0) == nil {
		return User{}, errors.New("ada yang salah")
	} else {
		user := argument.Get(0).(User)
		return user, nil
	}
}

func (r *RepositoryMock) FindByEmail(email string) (User, error) {
	var user User
	argument := r.Mock.Called(email)
	if argument.Get(0) == nil {
		return user, errors.New("ada yang salah")
	} else {
		user := argument.Get(0).(User)
		return user, nil
	}
}

func (r *RepositoryMock) Update(user User) (User, error) {
	argument := r.Mock.Called(user)
	if argument.Get(0) == nil {
		return user, errors.New("ada yang salah")
	} else {
		newUser := argument.Get(0).(User)
		return newUser, nil
	}
}
func (r *RepositoryMock) UpdateByID(user User) (User, error) {
	argument := r.Mock.Called(user)
	if argument.Get(0) == nil {
		return user, errors.New("ada yang salah")
	} else {
		newUser := argument.Get(0).(User)
		return newUser, nil
	}
}

func (r *RepositoryMock) DeleteUser(user User) (User, error) {
	argument := r.Mock.Called(user)
	if argument.Get(0) == nil {
		return user, errors.New("ada yang salah")
	} else {
		newUser := argument.Get(0).(User)
		return newUser, nil
	}

}

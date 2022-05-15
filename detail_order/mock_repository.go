package detail

import (
	"errors"
	"order_kafe/order"

	"github.com/stretchr/testify/mock"
)

type RepositoryMock struct {
	Mock mock.Mock
}

func NewRepositoryMock(mock mock.Mock) *RepositoryMock {
	return &RepositoryMock{mock}
}

func (r *RepositoryMock) Save(detail []order.DetailOrder) ([]order.DetailOrder, error) {
	argument := r.Mock.Called()
	if argument.Get(0) == nil {
		return []order.DetailOrder{}, errors.New("ada yang salah")
	} else {
		category := argument.Get(0).([]order.DetailOrder)
		return category, nil
	}
}

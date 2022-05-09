package order

import "order_kafe/user"

type OrderService interface {
	CreateOrder(input InputNewOrder) (Order, error)
}

type orderService struct {
	repository OrderRepository
	userRepo   user.UserRepository
}

func NewOrderService(repository OrderRepository, userRepo user.UserRepository) *orderService {
	return &orderService{repository, userRepo}
}

func (s *orderService) CreateOrder(input InputNewOrder) (Order, error) {
	var order Order

	order.Infomation = "payment checking"

	user, errUser := s.userRepo.FindById(input.UserID)
	if errUser != nil {
		return order, errUser
	}
	order.UserID = user.ID
	order.User = user

	newOrder, errOrder := s.repository.Save(order)
	if errOrder != nil {
		return order, errOrder
	}

	return newOrder, nil
}

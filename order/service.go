package order

import "order_kafe/user"

type OrderService interface {
	CreateOrder(input InputNewOrder) (Order, error)
}

type orderService struct {
	repository     orderRepository
	userRepository user.UserRepository
}

func NewOrderService(repository orderRepository, userRepository user.UserRepository) *orderService {
	return &orderService{repository, userRepository}
}

func (s *orderService) CreateOrder(input InputNewOrder) (Order, error) {
	var order Order

	order.Infomation = "payment checking"

	user, errUser := s.userRepository.FindById(input.UserID)
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

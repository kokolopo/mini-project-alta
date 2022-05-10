package order

import "order_kafe/user"

type OrderService interface {
	CreateOrder(id int) (Order, error)
}

type orderService struct {
	repository OrderRepository
	userRepo   user.UserRepository
}

func NewOrderService(repository OrderRepository, userRepo user.UserRepository) *orderService {
	return &orderService{repository, userRepo}
}

func (s *orderService) CreateOrder(id int) (Order, error) {
	var order Order

	user, errUser := s.userRepo.FindById(id)
	if errUser != nil {
		return order, errUser
	}
	order.UserID = user.ID
	order.User = user
	order.Infomation = "payment checking"

	newOrder, errOrder := s.repository.Save(order)
	if errOrder != nil {
		return order, errOrder
	}

	return newOrder, nil
}

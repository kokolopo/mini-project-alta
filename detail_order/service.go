package detail

import "order_kafe/order"

type DetailOrderService interface {
	SaveDetailOrder(orderId int, input []InputNewDetailOrder) ([]order.DetailOrder, error)
}

type detailOrderService struct {
	repository DetailOrderRepository
}

func NewDetailOrderService(repository DetailOrderRepository) *detailOrderService {
	return &detailOrderService{repository}
}

func (s *detailOrderService) SaveDetailOrder(orderId int, input []InputNewDetailOrder) ([]order.DetailOrder, error) {
	var detail []order.DetailOrder

	for _, v := range input {
		detail = append(detail, order.DetailOrder{OrderID: orderId, ItemID: v.ItemID, Quantity: v.Quantity, Note: v.Note})
	}

	//save data yang sudah dimapping kedalam struct DetailOrder
	newDetail, err := s.repository.Save(detail)
	if err != nil {
		return newDetail, err
	}

	return newDetail, nil
}

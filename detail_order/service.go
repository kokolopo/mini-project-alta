package detailorder

type DetailOrderService interface {
	SaveDetailOrder(input []InputNewDetailOrder) ([]DetailOrder, error)
}

type detailOrderService struct {
	repository detailOrderRepository
}

func NewDetailOrderService(repository detailOrderRepository) *detailOrderService {
	return &detailOrderService{repository}
}

func (s *detailOrderService) SaveDetailOrder(input []InputNewDetailOrder) ([]DetailOrder, error) {
	var detail []DetailOrder

	for _, v := range input {
		detail = append(detail, DetailOrder{OrderID: v.OrderID, ItemID: v.ItemID, Quantity: v.Quantity, Note: v.Note})
	}

	details, err := s.repository.Save(detail)
	if err != nil {
		return details, err
	}

	return details, nil
}

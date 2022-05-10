package detail

type DetailOrderService interface {
	SaveDetailOrder(input []InputNewDetailOrder) ([]DetailOrder, error)
}

type detailOrderService struct {
	repository DetailOrderRepository
}

func NewDetailOrderService(repository DetailOrderRepository) *detailOrderService {
	return &detailOrderService{repository}
}

func (s *detailOrderService) SaveDetailOrder(input []InputNewDetailOrder) ([]DetailOrder, error) {
	var detail []DetailOrder

	//tangkap nilai dari inputan
	// detail.OrderID = input.OrderID
	// detail.ItemID = input.ItemID
	// detail.Quantity = input.Quantity
	// detail.Note = input.Note

	for _, v := range input {
		detail = append(detail, DetailOrder{OrderID: v.OrderID, ItemID: v.ItemID, Quantity: v.Quantity, Note: v.Note})
	}

	// for _, v := range detail {
	// 	v.Order.UserID = 1 //didapatkan dari user yng sedang login
	// }

	//save data yang sudah dimapping kedalam struct DetailOrder
	newDetail, err := s.repository.Save(detail)
	if err != nil {
		return newDetail, err
	}

	return newDetail, nil
}

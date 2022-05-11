package transaction

type TransactionService interface {
	GetTransactionByUserId(userId int) ([]Transaction, error)
}

type transactionService struct {
	repository TransactionRepository
	// userRepo   user.UserRepository
	// orderRepo  order.OrderRepository
}

func NewTransactionService(repository TransactionRepository) *transactionService {
	return &transactionService{repository}
}

func (s *transactionService) GetTransactionByUserId(userId int) ([]Transaction, error) {
	transactions, err := s.repository.GetByUserId(userId)
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

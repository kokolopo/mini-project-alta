package transaction

import (
	"order_kafe/order"
	"order_kafe/payment"
	"order_kafe/user"
	"strconv"
)

type TransactionService interface {
	GetTransactionByUserId(userId int) ([]Transaction, error)
	GetTransactions() ([]Transaction, error)
	CreateTransaction(input InputCreateTansaction) (Transaction, error)
	ProcessPayment(input InputTransactionNotif) error
}

type transactionService struct {
	repository     TransactionRepository
	paymentService payment.Service
	userRepo       user.UserRepository
	orderRepo      order.OrderRepository
}

func NewTransactionService(repository TransactionRepository, paymentService payment.Service, userRepo user.UserRepository, orderRepo order.OrderRepository) *transactionService {
	return &transactionService{repository, paymentService, userRepo, orderRepo}
}

func (s *transactionService) GetTransactionByUserId(userId int) ([]Transaction, error) {
	transactions, err := s.repository.GetByUserId(userId)
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (s *transactionService) GetTransactions() ([]Transaction, error) {
	transactions, err := s.repository.GetTransactions()
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (s *transactionService) CreateTransaction(input InputCreateTansaction) (Transaction, error) {
	transaction := Transaction{}
	transaction.UserID = input.User.ID
	transaction.OrderID = input.OrderID
	transaction.Amount = input.Amount
	transaction.Status = "pending"
	transaction.Code = "ODR" + strconv.Itoa(input.OrderID) + "-" + strconv.Itoa(input.User.ID)

	newTrans, err := s.repository.Save(transaction)
	if err != nil {
		return newTrans, err
	}

	paymentTransaction := payment.Transaction{
		ID:     newTrans.ID,
		Amount: newTrans.Amount,
	}

	paymentUrl, errUrl := s.paymentService.GetPaymentUrl(paymentTransaction, input.User)
	if errUrl != nil {
		return newTrans, err
	}
	newTrans.PaymentUrl = paymentUrl

	updateTrans, errUpdate := s.repository.Update(newTrans)
	if errUpdate != nil {
		return updateTrans, errUpdate
	}

	return newTrans, nil
}

func (s *transactionService) ProcessPayment(input InputTransactionNotif) error {
	transaction_id, _ := strconv.Atoi(input.OrderID)

	transaction, err := s.repository.GetById(transaction_id)
	if err != nil {
		return err
	}

	if input.PaymentType == "credit_card" && input.TransactionStatus == "capture" && input.FraudStatus == "accept" {
		transaction.Status = "paid"
	} else if input.TransactionStatus == "settlement" {
		transaction.Status = "paid"
	} else if input.TransactionStatus == "deny" || input.TransactionStatus == "expire" || input.TransactionStatus == "cancel" {
		transaction.Status = "cancelled"
	}

	updatedTransaction, err := s.repository.Update(transaction)
	if err != nil {
		return err
	}

	campaign, err := s.orderRepo.FindById(updatedTransaction.OrderID)
	if err != nil {
		return err
	}

	if updatedTransaction.Status == "paid" {
		campaign.Infomation = "PAID AND MUST BE MADE"

		_, err := s.orderRepo.Update(campaign)
		if err != nil {
			return err
		}
	}

	return nil
}

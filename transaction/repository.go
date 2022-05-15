package transaction

import "gorm.io/gorm"

type TransactionRepository interface {
	GetByUserId(userId int) ([]Transaction, error)
	GetTransactions() ([]Transaction, error)
	Save(transaction Transaction) (Transaction, error)
	Update(transaction Transaction) (Transaction, error)
}

type transactionRepository struct {
	DB *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *transactionRepository {
	return &transactionRepository{db}
}

func (r *transactionRepository) GetByUserId(userId int) ([]Transaction, error) {
	var transactions []Transaction

	err := r.DB.Preload("User").Where("user_id = ?", userId).Order("id desc").Find(&transactions).Error
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (r *transactionRepository) GetTransactions() ([]Transaction, error) {
	var transactions []Transaction

	err := r.DB.Preload("User").Order("id desc").Find(&transactions).Error
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (r *transactionRepository) Save(transaction Transaction) (Transaction, error) {
	err := r.DB.Create(&transaction).Error
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (r *transactionRepository) Update(transaction Transaction) (Transaction, error) {
	err := r.DB.Save(&transaction).Error
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

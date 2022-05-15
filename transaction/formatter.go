package transaction

import "time"

type UserTransactionFormatter struct {
	ID        int       `json:"id"`
	OrderID   int       `json:"order_id"`
	Fullname  string    `json:"fullname"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

func FormatUserTransaction(transaction Transaction) UserTransactionFormatter {
	formatter := UserTransactionFormatter{}

	formatter.ID = transaction.ID
	formatter.OrderID = transaction.OrderID
	formatter.Fullname = transaction.User.Fullname
	formatter.Amount = transaction.Amount
	formatter.CreatedAt = transaction.CreatedAt

	return formatter
}

func FormatUserTransactions(transactions []Transaction) []UserTransactionFormatter {
	if len(transactions) == 0 {
		return []UserTransactionFormatter{}
	}

	var transactionsFormatter []UserTransactionFormatter

	for _, transaction := range transactions {
		formatter := FormatUserTransaction(transaction)
		transactionsFormatter = append(transactionsFormatter, formatter)
	}

	return transactionsFormatter
}

type TransactionFormatter struct {
	ID         int       `json:"id"`
	OrderID    int       `json:"order_id"`
	UserID     int       `json:"user_id"`
	Amount     int       `json:"amount"`
	Status     string    `json:"status"`
	Code       string    `json:"code"`
	PaymentUrl string    `json:"payment_url"`
	CreatedAt  time.Time `json:"created_at"`
}

func FormatTransaction(transaction Transaction) TransactionFormatter {
	formatter := TransactionFormatter{}

	formatter.ID = transaction.ID
	formatter.OrderID = transaction.OrderID
	formatter.UserID = transaction.UserID
	formatter.Amount = transaction.Amount
	formatter.Status = transaction.Status
	formatter.Code = transaction.Code
	formatter.PaymentUrl = transaction.PaymentUrl
	formatter.CreatedAt = transaction.CreatedAt

	return formatter
}

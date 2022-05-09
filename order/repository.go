package order

import "gorm.io/gorm"

type OrderRepository interface {
	Save(order Order) (Order, error)
	FetchAll() ([]Order, error)
	FindById(id int) (Order, error)
	Update(order Order) (Order, error)
	Delete(order Order) (Order, error)
}

type orderRepository struct {
	DB *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *orderRepository {
	return &orderRepository{db}
}

func (r *orderRepository) Save(order Order) (Order, error) {
	err := r.DB.Create(&order).Error
	if err != nil {
		return order, err
	}

	return order, nil
}
package order

import "gorm.io/gorm"

type OrderRepository interface {
	Save(order Order) (Order, error)
	FetchAll() ([]Order, error)
	FindById(id int) (Order, error)
	Update(order Order) (Order, error)
	// FindByEmail(email string) (Order, error)
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

func (r *orderRepository) FetchAll() ([]Order, error) {
	var orders []Order

	err := r.DB.Preload("Details", "detail_orders.order_id").Preload("Details.Item").Preload("User").Order("id desc").Find(&orders).Error
	if err != nil {
		return orders, err
	}

	return orders, nil
}

func (r *orderRepository) FindById(id int) (Order, error) {
	var order Order

	err := r.DB.Preload("User").Where("id = ?", id).Find(&order).Error

	if err != nil {
		return order, err
	}

	return order, nil
}

func (r *orderRepository) Update(order Order) (Order, error) {
	err := r.DB.Save(&order).Error

	if err != nil {
		return order, err
	}

	return order, nil
}

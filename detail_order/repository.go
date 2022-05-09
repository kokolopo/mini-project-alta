package detailorder

import "gorm.io/gorm"

type DetailOrderRepository interface {
	Save(detailOrder []DetailOrder) (DetailOrder, error)
	// FetchAll() ([]DetailOrder, error)
	// FindById(id int) (DetailOrder, error)
	// Update(detailOrder DetailOrder) (DetailOrder, error)
	// Delete(detailOrder DetailOrder) (DetailOrder, error)
}

type detailOrderRepository struct {
	DB *gorm.DB
}

func NewDetailOrderRepository(db *gorm.DB) *detailOrderRepository {
	return &detailOrderRepository{db}
}

func (r *detailOrderRepository) Save(detailOrder []DetailOrder) ([]DetailOrder, error) {
	err := r.DB.Create(&detailOrder).Error
	if err != nil {
		return detailOrder, err
	}

	return detailOrder, nil
}

package detail

import "gorm.io/gorm"

type DetailOrderRepository interface {
	Save(detail []DetailOrder) ([]DetailOrder, error)
}

type detailOrderRepository struct {
	DB *gorm.DB
}

func NewDetailOrderRepository(db *gorm.DB) *detailOrderRepository {
	return &detailOrderRepository{db}
}

func (r *detailOrderRepository) Save(detail []DetailOrder) ([]DetailOrder, error) {
	err := r.DB.Create(&detail).Error
	if err != nil {
		return detail, err
	}

	return detail, nil
}

package repositories

import (
	"example.com/web-service-gin/src/models"
	"gorm.io/gorm"
)

type OrderRepository struct {
	DB *gorm.DB
}

func (r *OrderRepository) Create(payload *models.Order) (models.Order, error) {
	result := r.DB.Create(payload)

	return *payload, result.Error
}

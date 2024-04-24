package services

import (
	"example.com/web-service-gin/src/dtos"
	"example.com/web-service-gin/src/models"
	"example.com/web-service-gin/src/repositories"
)

type OrderService struct {
	Repo *repositories.OrderRepository
}

func (s *OrderService) Create(payload *models.Order) (models.Order, error) {
	return s.Repo.Create(payload)
}

func (s *OrderService) MapCreateOrderDtoToModel(dto *dtos.CreateOrderDto, point float64) *models.Order {

	// Initialize the model payload
	order := &models.Order{
		Name:   dto.Name,
		Email:  dto.Email,
		BookId: dto.BookId,
		Point:  point,
		Status: "Pending", // Set default status
	}

	return order
}

package services_test

import (
	"testing"

	"example.com/web-service-gin/configs"
	"example.com/web-service-gin/src/models"
	"example.com/web-service-gin/src/repositories"
	"example.com/web-service-gin/src/services"
	"gorm.io/gorm"

	"github.com/stretchr/testify/assert"
)

func NewOrderRepository(DB *gorm.DB) *repositories.OrderRepository {
	return &repositories.OrderRepository{
		DB: DB,
	}
}

func TestOrderService_Create(t *testing.T) {
	// Create a mock repository
	sqlDB, db, _ := configs.DbMock(t)
	defer sqlDB.Close()
	repo := NewOrderRepository(db)

	// Create the service with the mock repository
	service := services.OrderService{Repo: repo}

	// Prepare the input data
	payload := &models.Order{
		Name:   "John Doe",
		Email:  "john@example.com",
		BookId: 1,
		Point:  10,
		Status: "Pending",
	}

	// Call the Create method of the OrderService
	createdOrder, _ := service.Create(payload)

	// // Assert that the created order matches the input payload
	assert.NotNil(t, createdOrder)
	assert.Equal(t, payload.Name, createdOrder.Name)
	assert.Equal(t, payload.Email, createdOrder.Email)
	assert.Equal(t, payload.BookId, createdOrder.BookId)
	assert.Equal(t, payload.Point, createdOrder.Point)
	assert.Equal(t, payload.Status, createdOrder.Status)
}

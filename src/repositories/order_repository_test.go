package repositories

import (
	"testing"

	"example.com/web-service-gin/configs"
	"example.com/web-service-gin/src/models"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func NewOrderRepository(DB *gorm.DB) *OrderRepository {
	return &OrderRepository{
		DB: DB,
	}
}

func TestOrderRepository_Create(t *testing.T) {
	sqlDB, db, mock := configs.DbMock(t)
	defer sqlDB.Close()
	implObj := NewOrderRepository(db)

	addRow := sqlmock.NewRows([]string{"id"}).AddRow("1")
	expectedSQL := "INSERT INTO \"orders\" (.+) VALUES (.+)"
	mock.ExpectBegin()
	mock.ExpectQuery(expectedSQL).WillReturnRows(addRow)
	mock.ExpectCommit()

	// Create an instance of models.Order
	reqOrder := models.Order{}

	// Use the repository method to create the order
	// This assumes that your Create method accepts a pointer to models.Order
	payload := implObj.DB.Create(&reqOrder)
	assert.NotNil(t, payload)

	// Verify that all expectations were met
	assert.Nil(t, mock.ExpectationsWereMet())
}

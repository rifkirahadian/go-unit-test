package repositories

import (
	"testing"

	"example.com/web-service-gin/configs"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func NewBookRepository(DB *gorm.DB) *BookRepository {
	return &BookRepository{
		DB: DB,
	}
}

func TestBookRepository_GetAll(t *testing.T) {
	sqlDB, db, mock := configs.DbMock(t)
	defer sqlDB.Close()

	implObj := NewBookRepository(db)
	books := sqlmock.NewRows([]string{"id", "title", "writer", "cover_image"}).
		AddRow(1, "Clean Code", "John Doe", "https:/image.jpg")

	expectedSQL := "SELECT (.+) FROM \"books\""
	mock.ExpectQuery(expectedSQL).WillReturnRows(books)
	_, res := implObj.GetAll()
	assert.Nil(t, res)
	assert.Nil(t, mock.ExpectationsWereMet())
}

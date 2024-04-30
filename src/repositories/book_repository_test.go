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
		AddRow(1, "Clean Code", "John Doe", "https:/image1.jpg").
		AddRow(1, "Dirty Code", "John Cena", "https:/image2.jpg")

	expectedSQL := "SELECT (.+) FROM \"books\""
	mock.ExpectQuery(expectedSQL).WillReturnRows(books)
	res, err := implObj.GetAll()
	assert.Nil(t, err)
	assert.Equal(t, len(res), 2)
	assert.Equal(t, res[0].Title, "Clean Code")
	assert.Equal(t, res[0].Writer, "John Doe")
	assert.Equal(t, res[0].CoverImage, "https:/image1.jpg")

	assert.Nil(t, mock.ExpectationsWereMet())
}

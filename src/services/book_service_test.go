package services_test

import (
	"testing"

	"example.com/web-service-gin/configs"
	"example.com/web-service-gin/src/repositories"
	"example.com/web-service-gin/src/services"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func NewBookRepository(DB *gorm.DB) *repositories.BookRepository {
	return &repositories.BookRepository{
		DB: DB,
	}
}

func TestBookService_GetAllBook(t *testing.T) {
	// Create a mock repository
	sqlDB, db, mock := configs.DbMock(t)
	defer sqlDB.Close()
	repo := NewBookRepository(db)

	// Create the service with the mock repository
	service := services.BookService{Repo: repo}

	books := sqlmock.NewRows([]string{"id", "title", "writer", "cover_image"}).
		AddRow(1, "Clean Code", "John Doe", "https:/image1.jpg").
		AddRow(1, "Dirty Code", "John Cena", "https:/image2.jpg")

	expectedSQL := "SELECT (.+) FROM \"books\""
	mock.ExpectQuery(expectedSQL).WillReturnRows(books)

	res, err := service.GetAllBook()
	assert.Nil(t, err)
	assert.Equal(t, len(res), 2)
	assert.Equal(t, res[0].Title, "Clean Code")
	assert.Equal(t, res[0].Writer, "John Doe")
	assert.Equal(t, res[0].CoverImage, "https:/image1.jpg")

	assert.Nil(t, mock.ExpectationsWereMet())
}

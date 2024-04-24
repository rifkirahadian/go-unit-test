package repositories

import (
	"example.com/web-service-gin/src/models"
	"gorm.io/gorm"
)

type BookRepository struct {
	DB *gorm.DB
}

func (r *BookRepository) GetAll() ([]models.Book, error) {
	var books []models.Book
	if err := r.DB.Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

func (r *BookRepository) FindOne(id int16) (models.Book, error) {
	var book models.Book
	result := r.DB.First(&book, id)

	return book, result.Error
}

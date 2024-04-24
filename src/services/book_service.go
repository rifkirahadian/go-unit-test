package services

import (
	"example.com/web-service-gin/src/models"
	"example.com/web-service-gin/src/repositories"
)

type BookService struct {
	Repo *repositories.BookRepository
}

func (s *BookService) GetAllBook() ([]models.Book, error) {
	return s.Repo.GetAll()
}

func (s *BookService) FindOneBook(id int16) (models.Book, error) {
	return s.Repo.FindOne(id)
}

package main

import (
	"example.com/web-service-gin/src/controllers"
	"example.com/web-service-gin/src/repositories"
	"example.com/web-service-gin/src/services"
	"gorm.io/gorm"
)

type Container struct {
	BookController *controllers.BookController
}

func NewContainer(db *gorm.DB) *Container {
	// Initialize repositories
	bookRepo := &repositories.BookRepository{DB: db}

	// Initialize services
	bookService := &services.BookService{Repo: bookRepo}

	// Initialize controllers
	bookController := &controllers.BookController{Service: bookService}

	return &Container{
		BookController: bookController,
	}
}

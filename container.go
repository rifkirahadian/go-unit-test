package main

import (
	"example.com/web-service-gin/src/controllers"
	"example.com/web-service-gin/src/repositories"
	"example.com/web-service-gin/src/services"
	"gorm.io/gorm"
)

type Container struct {
	BookController  *controllers.BookController
	OrderController *controllers.OrderController
}

func NewContainer(db *gorm.DB) *Container {
	// Initialize repositories
	bookRepo := &repositories.BookRepository{DB: db}
	orderRepo := &repositories.OrderRepository{DB: db}

	// Initialize services
	bookService := &services.BookService{Repo: bookRepo}
	orderService := &services.OrderService{Repo: orderRepo}

	// Initialize controllers
	bookController := &controllers.BookController{Service: bookService}
	orderController := &controllers.OrderController{Service: orderService, BookService: bookService}

	return &Container{
		BookController:  bookController,
		OrderController: orderController,
	}
}

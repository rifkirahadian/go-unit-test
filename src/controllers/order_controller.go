package controllers

import (
	"net/http"

	"example.com/web-service-gin/src/dtos"
	"example.com/web-service-gin/src/services"
	"example.com/web-service-gin/src/utils"
	"github.com/gin-gonic/gin"
)

type OrderController struct {
	Service     *services.OrderService
	BookService *services.BookService
}

func (c *OrderController) CreateBook(ctx *gin.Context) {
	var body dtos.CreateOrderDto

	if err := ctx.ShouldBindJSON(&body); err != nil {
		utils.ValidationResponse(ctx, err)
		return
	}

	book, _ := c.BookService.FindOneBook(body.BookId)
	if book.ID == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	payload := c.Service.MapCreateOrderDtoToModel(&body, book.Point)

	order, _ := c.Service.Create(payload)

	ctx.IndentedJSON(http.StatusCreated, order)
}

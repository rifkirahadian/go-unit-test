package controllers

import (
	"net/http"

	"example.com/web-service-gin/src/services"
	"github.com/gin-gonic/gin"
)

type BookController struct {
	Service *services.BookService
}

func (c *BookController) FindBooks(ctx *gin.Context) {
	book, err := c.Service.GetAllBook()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.IndentedJSON(http.StatusCreated, book)
}

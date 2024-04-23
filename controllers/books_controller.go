package controllers

import (
	"net/http"

	"example.com/web-service-gin/configs"
	"example.com/web-service-gin/models"
	"github.com/gin-gonic/gin"
)

func FindBooks(c *gin.Context) {
	db := configs.InitDB()
	var book []models.Book
	db.Find(&book)

	c.IndentedJSON(http.StatusCreated, book)
}

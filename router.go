package main

import (
	"example.com/web-service-gin/src/controllers"
	"github.com/gin-gonic/gin"
)

func GetRouter(router *gin.Engine, container *Container) {
	router.GET("/albums", controllers.GetAlbums)
	router.POST("/albums", controllers.PostAlbums)
	router.GET("/albums/:id", controllers.GetAlbumByID)

	router.GET("/books", container.BookController.FindBooks)
}

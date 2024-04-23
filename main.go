package main

import (
	"example.com/web-service-gin/configs"
	"github.com/gin-gonic/gin"
)

func main() {
	db := configs.InitDB()

	router := gin.Default()
	container := NewContainer(db)
	GetRouter(router, container)

	router.Run("localhost:8080")
}

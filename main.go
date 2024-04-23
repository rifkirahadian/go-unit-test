package main

import (
	"example.com/web-service-gin/configs"
	"github.com/gin-gonic/gin"
)

func main() {
	configs.InitDB()

	router := gin.Default()
	GetRouter(router)

	router.Run("localhost:8080")
}

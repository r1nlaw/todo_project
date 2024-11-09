package main

import (
	"todo_project/server.go"

	"github.com/gin-gonic/gin"
)

func init() {
	server.InitServer() // Инициализируем сервер
}

func main() {
	server.StartServer() // Запускаем сервер
	// Создаем новый роутер по умолчанию в GIN
	router := gin.Default()
	// Регистрируем обработчик GET-запроса /ping
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Привет я GIN",
		})

	})
	// Запускаем роутер
	router.Run(":8080")
}

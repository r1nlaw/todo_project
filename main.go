package main

import "github.com/gin-gonic/gin"

func main() {
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

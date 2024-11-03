package server

import "github.com/gin-gonic/gin"

func InitRoutes() {
	router := gin.Default()    // Инициализация роута (по умолчанию)
	router.PUT("/note/:id")    // Редактирование заметки по id
	router.DELETE("/note/:id") // Удаление заметки по id
	router.GET("/note/:id")    // Получение заметки по id
	router.POST("/note")       // Создание заметки
	router.GET("/notes")       // Получение списка всех заметок

	// Запуск сервера на порту 8080
	router.Run(":8080")
}

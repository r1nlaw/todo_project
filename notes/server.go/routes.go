package server

import (
	"todo_project/handlers"

	"github.com/gin-gonic/gin"
)

func InitRoutes() {
	router := gin.Default()                                // Инициализация роута (по умолчанию)
	router.PUT("/note/:id", handlers.UpdateNoteHandler)    // Редактирование заметки по id
	router.DELETE("/note/:id", handlers.DeleteNoteHandler) // Удаление заметки по id
	router.GET("/note/:id", handlers.GetNoteHandler)       // Получение заметки по id
	router.POST("/note", handlers.CreateNoteHandler)       // Создание заметки
	router.GET("/notes", handlers.GetAllNotesHandler)      // Получение списка всех заметок

	// Запуск сервера на порту 8080
	router.Run(":8080")
}

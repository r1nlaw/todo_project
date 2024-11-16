package server

import (
	"auth/handlers"

	"github.com/gin-gonic/gin"
)

func InitRoutes() {
	// Инициализация роута (по умолчанию)
	router := gin.Default()
	// Создание пользователя
	router.PUT("/user", handlers.RegisterUserHandler)
	// Авторизация пользователя
	router.POST("/user", handlers.SignInHandler)
	// Обновление токена
	router.PUT("/refresh", handlers.RefreshTokenHandler)
	// Получение данных пользователя
	router.GET("/user", handlers.GetUserHandler)

	auth := router.Group("/")
	auth.Use(handlers.AuthMiddleware())
	{
		// Получение данных от пользователя, если пропустит перехватчик
		router.GET("/user", handlers.GetUserHandler)
	}
	router.Run(":9104")
}

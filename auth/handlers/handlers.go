package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignInHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Авторизация прошла успешно"})
}
func RefreshTokenHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Токен обновлен"})
}
func GetUserHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Получение данных пользователя прошло успешно"})
}

func RegisterUserHandler(ctx *gin.Context) {
	registerUserHandler(ctx)
}

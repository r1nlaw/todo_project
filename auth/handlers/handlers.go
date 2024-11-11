package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignInHandler(ctx *gin.Context) {
	signIn(ctx)
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

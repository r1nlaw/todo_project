package utils

import (
	"auth/envs"
	"auth/models"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

// Хэшируем пароль при помощи bcrypt
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10) // 10 - стоимость хэширования (1024 итераций хэширования пароля)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// Проверяем хэш пароль при помощи bcrypt
func CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// Функция для генерации JWT-токена
func GenerateTokens(userID uint) (models.Tokens, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // Срок действия токена 24 часа (exp)
	})

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 600).Unix(), // Срок действия токена 1 месяц
	})
	signedAccessToken, _ := accessToken.SignedString([]byte(envs.ServerEnvs.JWT_SECRET))

	signedRefreshToken, _ := refreshToken.SignedString([]byte(envs.ServerEnvs.JWT_SECRET))

	return models.Tokens{AccessToken: signedAccessToken, RefreshToken: signedRefreshToken}, nil
}

// Функция проверки JWT-токена

func ValidateRefreshToken(tokenString string) (uint, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("неожиданный метод подписи: %v", token.Header["alg"])
		}
		return []byte(envs.ServerEnvs.JWT_SECRET), nil
	})
	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userIDValue, ok := claims["user_id"].(float64) // приведение к float64
		if !ok {
			return 0, fmt.Errorf("user_id claims is not float64")
		}
		return uint(userIDValue), nil // Конвертация float64 в uint
	} else {
		return 0, fmt.Errorf("недействительный токен")
	}
}

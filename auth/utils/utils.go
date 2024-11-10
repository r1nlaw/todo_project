package utils

import "golang.org/x/crypto/bcrypt"

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

package database

import (
	"fmt"
	"todo_project/envs"

	"github.com/go-redis/redis"
)

var RedisClient *redis.Client

// Подключение к Redis
func InitRedis() error {
	// формирование URI для подключения к Redis
	redisUri := fmt.Sprintf("%s:%s", envs.ServerEnvs.REDIS_HOST, envs.ServerEnvs.REDIS_PORT)

	// Создаем новый клиента Redis и присваивания его глобаной
	// переменной [database.RedisClient]
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     redisUri,
		Password: "", // Пароль не нужен
		DB:       0,  // База данных 0 (по умолчанию)
	})
	// Проверка соединения командой PING
	// в ответпри успешном соединении нам придет PONG или ошибка
	status := RedisClient.Ping()
	if status.Val() == "PONG" {
		return nil
	} else {
		return fmt.Errorf("Ошибка при подключении к Redis: %v", status)
	}
}

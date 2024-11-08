package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"todo_project/database"
	"todo_project/models"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

// Обработка запроса для получения заметки по ID
func GetNoteHandler(ctx *gin.Context) {
	authorId := 1
	// Получаем ID заметки из параметра запроса
	id := ctx.Param("id")
	// Получаем коллекцию "notes"
	collection := database.MongoClient.Database("admin").Collection(fmt.Sprintf("notes/%d", authorId))

	// Объявляем переменную для хранения заметки
	var note models.Note
	// Создаем фильтр для поиска по ID
	filter := bson.M{"id": id}
	// Ищем заметку в коллекции,
	// если она есть возращаем ее иначе сообщение об ошибке
	errFind := collection.FindOne(ctx, filter).Decode(&note)
	if errFind != nil {
		ctx.JSON(http.StatusOK, "Заметка не найдена")
	}
	// Возращаем заметку
	ctx.JSON(http.StatusOK, &note)

}

func GetAllNotesHandler(ctx *gin.Context) {
	authorId := 1
	// Объявляем список заметок
	var notes []models.Note

	// Проверяем, есть ли в кэше данные
	val, err := database.RedisClient.Get(fmt.Sprintf("notes/%d", authorId)).Result()
	if err == redis.Nil {
		// Получаем коллекцию "notes"
		collection := database.MongoClient.Database("admin").Collection(fmt.Sprintf("notes/%d", authorId))

		// Поиск документов без фильтров для получения всех заметок
		cursor, err := collection.Find(ctx, bson.M{})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		// Закрытие курсора, при завершении работы функции
		defer cursor.Close(ctx)
		// Итерация по курсору и декодирование документов в заметки
		for cursor.Next(ctx) {
			var note models.Note
			err := cursor.Decode(&note)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"error": err.Error(),
				})
				return
			}
			notes = append(notes, note)
		}
		// Проверка на ошибки после итерации
		if err := cursor.Err(); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		// Проверка на наличие заметок
		if len(notes) == 0 {
			ctx.JSON(http.StatusOK, "Заметок не найдено")

		} else {
			// Запись в кэш
			recordCacheToRedis(notes, authorId)
			// Возвращаем список заметок
			ctx.JSON(http.StatusOK, notes)
		}

	} else {
		getFormCache(val, ctx)
	}

}
func recordCacheToRedis(notes []models.Note, authorId int) {
	// Сериализуем список заметок в JSON
	notesJSON, err := json.Marshal(notes)
	// Обрабатываем ошибку или продолжаем без кэширования
	if err != nil {
		log.Printf("Ошибка при сериализации заметок: %v", err)
	} else {
		// Сохраняем сериализованные данные в Redis
		// Срок действия ключа - 30 минут
		err := database.RedisClient.Set(fmt.Sprintf("notes/%d", authorId), string(notesJSON), 1440*time.Minute).Err()
		// Обрабатываем ошибку или продолжаем без кэширования
		if err != nil {
			log.Printf("Ошибка при сохранении в Redis: %v", err)
		}
	}
}

func getFormCache(val string, ctx *gin.Context) {
	log.Printf("Кэш найден, загружаем из Кэша")
	notes := make([]models.Note, 0)
	json.Unmarshal([]byte(val), &notes)
	ctx.JSON(http.StatusOK, notes)
}

// Обработка запроса для удаления заметки по ID
func DeleteNoteHandler(ctx *gin.Context) {
	// Получаем ID заметки из параметра запроса
	id := ctx.Param("id")

	// Получаем коллекцию "notes"
	collection := database.MongoClient.Database("admin").Collection(fmt.Sprintf("notes/%d", 1))

	// Создаем фильтр для поиска по ID
	filter := bson.M{"id": id}

	// Удаляем заметку из коллекции по фильтру
	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	// Проверяем удалена ли заметка
	if result.DeletedCount == 0 {
		ctx.JSON(http.StatusOK, "Заметка не найдена")
	} else {
		ctx.JSON(http.StatusOK, "Заметка успешно удалена")
	}

}

// Обработка запроса для редактирования заметки по ID
func UpdateNoteHandler(ctx *gin.Context) {
	authorID := 1
	// Получаем ID заметки из параметра запроса
	id := ctx.Param("id")

	var note models.Note
	if err := ctx.ShouldBindJSON(&note); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Неверные данные",
		})
		return
	}

	// Получаем коллекцию "notes"
	collection := database.MongoClient.Database("admin").Collection(fmt.Sprintf("notes/%d", authorID))

	// Создаем динамический $set
	updateFields := bson.M{}
	// Проверяем, было ли передано имя заметки
	if note.Name != nil {
		updateFields["name"] = note.Name
	}
	// Проверяем, было ли передан контент заметки
	if note.Content != nil {
		updateFields["content"] = note.Content
	}
	// Создаем данные для обновления с помощью $set updateFields
	update := bson.M{"$set": updateFields}

	// Создаем фильтр для поиска по ID
	filter := bson.M{"id": id}

	// Обновляем заметку в коллекции по фильтру
	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Проверяем, обновлена ли заметка
	if result.MatchedCount == 0 {
		ctx.JSON(http.StatusOK, "Заметка не найдена")
	} else {
		ctx.JSON(http.StatusOK, "Заметка успешно отредактирована")
	}

}

// Обработка запроса для создания заметки
func CreateNoteHandler(ctx *gin.Context) {

	// Создание новой заметки
	var note models.Note
	// Получаем данные из запроса
	if err := ctx.ShouldBindJSON(&note); err != nil { // ctx.ShouldBindJSON - автоматически десериализует JSON и заполняет структуры Note
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Неверные данные",
		})
		return
	}

	// Получить уникальный id
	note.Id = uuid.New().String()
	// Тестовый id Автора
	note.AuthorID = 1

	// Получаем коллекцию "notes"
	collection := database.MongoClient.Database("admin").Collection(fmt.Sprintf("notes/%d", note.AuthorID))
	// Вставляем заметку в коллекцию
	_, errInsert := collection.InsertOne(ctx, note)
	if errInsert != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": errInsert.Error()})
	}
	// Если ошибок нет, то возвращаем заметку и статус 200
	ctx.JSON(http.StatusOK, gin.H{
		"note":    note,
		"message": "Заметка успешно создана",
	})

}

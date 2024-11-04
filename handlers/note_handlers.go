package handlers

import (
	"fmt"
	"net/http"
	"todo_project/database"
	"todo_project/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
)

func GetNoteHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "GetNoteHandler")
}

func GetAllNotesHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "GetAllNotesHandler")
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

func UpdateNoteHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "UpdateNoteHandler")
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

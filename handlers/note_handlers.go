package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetNoteHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "GetNoteHandler")
}

func GetAllNotesHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "GetAllNotesHandler")
}

func DeleteNoteHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "DeleteNoteHandler")
}

func UpdateNoteHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "UpdateNoteHandler")
}

func CreateNoteHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "CreateNoteHandler")
}

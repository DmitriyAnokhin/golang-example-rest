package controllers

import (
	"example/rest/models"
	"gorm.io/gorm"
	"net/http"

	"github.com/gin-gonic/gin"
)

var DB *gorm.DB

type CreateTrackInput struct {
	Artist string `json:"artist" binding:"required"`
	Title  string `json:"title" binding:"required"`
}

type UpdateTrackInput struct {
	Artist string `json:"artist"`
	Title  string `json:"title"`
}

func init() {
	DB = models.ConnectDB()
}

// GetAllTracks GET /tracks
// Получаем список всех треков
func GetAllTracks(context *gin.Context) {

	var tracks []models.Track

	DB.Find(&tracks)

	context.JSON(http.StatusOK, gin.H{"tracks": tracks})
}

// CreateTrack POST /tracks
// Создание трека
func CreateTrack(context *gin.Context) {

	var input CreateTrackInput

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	track := models.Track{Artist: input.Artist, Title: input.Title}

	DB.Create(&track)

	context.JSON(http.StatusOK, gin.H{"tracks": track})
}

// GetTrack GET /tracks/:id
// Получение одного трека по ID
func GetTrack(context *gin.Context) {

	// Проверяем имеется ли запись
	var track models.Track

	if err := DB.Where("id = ?", context.Param("id")).First(&track).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Запись не существует"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"tracks": track})
}

// UpdateTrack PATCH /tracks/:id
// Изменения информации
func UpdateTrack(context *gin.Context) {

	// Проверяем имеется ли такая запись перед тем как её менять
	var track models.Track

	if err := DB.Where("id = ?", context.Param("id")).First(&track).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Запись не существует"})
		return
	}

	var input UpdateTrackInput

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//models.ConnectDB().Model(&track).Update(input)

	track.Artist = input.Artist
	track.Title = input.Title

	DB.Save(&track)

	context.JSON(http.StatusOK, gin.H{"tracks": track})
}

// DeleteTrack DELETE /tracks/:id
// Удаление
func DeleteTrack(context *gin.Context) {

	// Проверяем имеется ли такая запись перед тем как её удалять
	var track models.Track

	if err := DB.Where("id = ?", context.Param("id")).First(&track).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Запись не существует"})
		return
	}

	DB.Delete(&track)

	context.JSON(http.StatusOK, gin.H{"tracks": true})
}

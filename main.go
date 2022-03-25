package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"

	// "example/rest/models"

	"example/rest/controllers"
)

func main() {

	// load .env file from given path
	// we keep it empty it will load .env from current directory
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// getting env variables SITE_TITLE and DB_HOST
	siteTitle := os.Getenv("DB_HOST")
	dbHost := os.Getenv("DB_PORT")

	fmt.Printf("godotenv : %s = %s \n", "Site Title", siteTitle)
	fmt.Printf("godotenv : %s = %s \n", "DB Host", dbHost)

	route := gin.Default()

	// models.ConnectDB()
	route.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"siteTitle": siteTitle,
			"dbHost":    dbHost,
		})
	})

	route.GET("/env", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"siteTitle": siteTitle,
			"dbHost":    dbHost,
		})
	})

	route.GET("/tracks", controllers.GetAllTracks)
	route.POST("/tracks", controllers.CreateTrack)
	route.GET("/tracks/:id", controllers.GetTrack)
	route.PATCH("/tracks/:id", controllers.UpdateTrack)
	route.DELETE("/tracks/:id", controllers.DeleteTrack)

	route.Run()
}

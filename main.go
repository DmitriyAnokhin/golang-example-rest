package main

import (
	"github.com/gin-gonic/gin"
	// "example/rest/models"

	"example/rest/controllers"
)

func main() {

	route := gin.Default()

	// models.ConnectDB()
	route.GET("/", controllers.Hello)

	route.GET("/tracks", controllers.GetAllTracks)
	route.POST("/tracks", controllers.CreateTrack)
	route.GET("/tracks/:id", controllers.GetTrack)
	route.PATCH("/tracks/:id", controllers.UpdateTrack)
	route.DELETE("/tracks/:id", controllers.DeleteTrack)

	route.Run()
}

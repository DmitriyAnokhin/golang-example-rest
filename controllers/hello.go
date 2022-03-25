package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Hello(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "Hello World!"})
}

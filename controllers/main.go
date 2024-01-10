package controllers

import (
	"github.com/Domains18/jwtauthsample/database"
	"github.com/Domains18/jwtauthsample/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterUser(context *gin.Context) {
	var user models.Users
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	if err := user.HashPassword(user.Password); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	record := database.Instance.Create(&user)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}
	context.JSON(http.StatusCreated, gin.H{"userID": user.ID, "email": user.Email, "username": user.Username})
}

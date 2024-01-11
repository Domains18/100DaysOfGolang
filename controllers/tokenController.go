package controllers

import (
	"github.com/Domains18/jwtauthsample/auth"
	"github.com/Domains18/jwtauthsample/database"
	"github.com/Domains18/jwtauthsample/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TokenRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func GenerateToken(context *gin.Context) {
	var request TokenRequest
	var user models.Users
	if err := context.ShouldBindJSON(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	record := database.Instance.Where("email= ?", request.Email).First(&user)
	if record.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}
	credentialErrors := user.CheckPassword(request.Password)
	if credentialErrors != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "invalid credentials"})
		context.Abort()
		return
	}
	tokenString, err := auth.JwtGenerator(user.Email, user.Username)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		context.Abort()
		return
	}
	context.JSON(http.StatusOK, gin.H{"token": tokenString})
}

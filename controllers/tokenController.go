package controllers

import "github.com/gin-gonic/gin"

type TokenRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

func GenerateToken(context *gin.Context){
	var request
}
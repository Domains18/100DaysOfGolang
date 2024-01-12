package main

import (
	"github.com/Domains18/jwtauthsample/controllers"
	"github.com/Domains18/jwtauthsample/database"
	"github.com/Domains18/jwtauthsample/middlewares"
	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect("connectionString")
	database.Migrate()

	router := initRouter()
	err := router.Run(":8080")
	if err != nil {
		return 
	}
}

func initRouter() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/token", controllers.GenerateToken)
		api.POST("/user/register", controllers.RegisterUser)
		secured := api.Group("/secured").Use(middlewares.Auth())
		{
			secured.GET("/ping", controllers.Ping)
		}
	}
	return router
}

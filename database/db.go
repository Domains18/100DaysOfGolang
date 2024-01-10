package database

import (
	"github.com/Domains18/jwtauthsample/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var Instance *gorm.DB
var databaseErrorExists error

type ErrorC struct {
	error
}

func Connect(connectionString string) {
	log.Println("starting connection string")
	Instance, databaseErrorExists = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if databaseErrorExists != nil {
		log.Fatal("database connection failure", databaseErrorExists)
		//panic("Cannot connect to database")
	}
	log.Println("Connected to database", Instance)
}


func Migrate(){
	err := Instance.AutoMigrate(&models.Users{})
	if err != nil {
		return 
	}
	log.Println("Database Migration Completed")
}
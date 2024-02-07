package database

import (
	"os"
	
	"gorm.io/gorm"
	"gorm.io/driver/mysql"

	"fiber-example/models"
)

var DB *gorm.DB

func Connect() {
	//DB_URI := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	DB_URI := os.Getenv("DB_URI")
	connection, err := gorm.Open(mysql.Open(DB_URI), &gorm.Config{})

	if err != nil {
		panic("could not connect to the database")
	}

	DB = connection

	connection.AutoMigrate(&models.User{})
}
package utils

import (
	"log"
	"os"
	
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"fiber-example/models"
)

func InitDB() {
	//DSN := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	DSN := os.Getenv("DB_URI")
	db, err := gorm.Open(mysql.Open(DSN), &gorm.Config{})

	DB = db

	if err != nil {
		log.Println(err)
		panic(err)
	}

	DB.AutoMigrate(&models.Todo{})
}
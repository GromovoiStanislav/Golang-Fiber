package database

import (
	"fmt"

	"gorm.io/gorm"
	"gorm.io/driver/sqlite"

	"fiber-example/internal/models"
)

var (
	DBConn *gorm.DB
)

func InitDatabase(dbName string) {
	var err error

	DBConn, err = gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Database connection successfully opened")

	DBConn.AutoMigrate(&models.Book{})
}
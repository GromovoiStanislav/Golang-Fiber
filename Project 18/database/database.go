package database

import (
	"fmt"

	"gorm.io/gorm"
	"gorm.io/driver/sqlite"

	"fiber-example/models"
)

var (
	DBConn *gorm.DB
)

func InitDatabase() {
	var err error
	DBConn, err = gorm.Open(sqlite.Open("books.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")


	DBConn.AutoMigrate(&models.Book{})
}
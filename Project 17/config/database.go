package config

import (
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"

	"fiber-example/models"
)

var (
	DB  *gorm.DB
	err error
)

func ManageMigrations() {
	DB, err = gorm.Open(sqlite.Open("blogs.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	
	err = DB.AutoMigrate(&models.Blogs{})
	if err != nil {
		panic("migration failure")
	}
}
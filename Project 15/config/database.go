package config

import (
    "os"
    
    "gorm.io/gorm"
	"gorm.io/driver/mysql"

    "fiber-example/entities"
)

var Database *gorm.DB

func Connect() error {
    var err error

	//DATABASE_URI := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	DATABASE_URI := os.Getenv("DATABASE_URI")

    Database, err = gorm.Open(mysql.Open(DATABASE_URI), &gorm.Config{
        SkipDefaultTransaction: true,
        PrepareStmt:            true,
    })

    if err != nil {
        panic(err)
    }

    Database.AutoMigrate(&entities.Dog{})

    return nil
}
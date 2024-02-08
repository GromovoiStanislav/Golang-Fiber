package models

import (
	"gorm.io/gorm"
)


type Book struct {
	gorm.Model
	ISIN   string `json:"ISIN" gorm:"unique"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}
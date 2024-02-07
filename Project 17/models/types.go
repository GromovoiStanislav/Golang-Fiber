package models

import "gorm.io/gorm"

type Message struct {
	Status      string
	Description string
}

type Blogs struct {
	gorm.Model
	Title  string `json:"name"`
	Author string `json:"author"`
	Reads  int    `json:"reads"`
}

type BlogDto struct {
	Name   string `json:"name" validate:"required"`
	Author string `json:"author" validate:"required"`
	Reads  int    `json:"reads" validate:"required"`
}


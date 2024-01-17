package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Book struct {
	//Id     string `json:"id" bson:"_id"`
	Id       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title  string `json:"title" bson:"title"`
	Author string `json:"author" bson:"author"`
	Year   string `json:"isbn" bson:"isbn"`
}
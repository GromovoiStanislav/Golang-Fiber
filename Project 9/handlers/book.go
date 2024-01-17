package handlers

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"fiber-example/database"
	"fiber-example/models"
)

type newBookDTO struct {
	Title     string `json:"title" bson:"title"`
	Author    string `json:"author" bson:"author"`
	Year      string `json:"year" bson:"year"`
}


func CreateBook(c *fiber.Ctx) error {
	getData := new(newBookDTO)
	if err := c.BodyParser(getData); err != nil {
		return err
	}


	createData := models.Book{
		Id:       primitive.NewObjectID(),
		Title:    getData.Title,
		Author:   getData.Author,
		Year:     getData.Year,
	}

	// get the id
	id := c.Params("id")
	if id == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "id is required",
		})
	}

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid id",
		})
	}

	// get the collection reference
	coll := database.GetCollection("libraries")

	// update one library
	_, err = coll.UpdateOne(context.TODO(), bson.M{"_id": objectId}, bson.D{{Key: "$push", Value: bson.M{"books": createData}}})

	// get the filter
	// filter := bson.D{{Key: "id", Value: createData.LibraryId}}
	// fmt.Println(filter)
	// nBookData := models.Book{
	// 	Title:  createData.Title,
	// 	Author: createData.Author,
	// 	ISBN:   createData.ISBN,
	// }
	// updatePayload := bson.D{{Key: "$push", Value: bson.M{"books": bson.M{"test": "test"}}}}

	// update the library
	if err != nil {
		return err
	}

	//fmt.Println(res.ModifiedCount)

	return c.SendString("Book created successfully")
}
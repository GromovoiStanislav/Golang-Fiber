package handlers

import (
	"context"
	"fmt"
	
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"fiber-example/database"
	"fiber-example/models"
)


// DELETE
func DeleteLibrary(c *fiber.Ctx) error {
	// get the id from the params
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

	libraryCollection := database.GetCollection("libraries")

	result, err := libraryCollection.DeleteOne(context.TODO(), bson.M{"_id": objectId})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error":   "Failed to delete library",
			"message": err.Error(),
		})
	}

	if result.DeletedCount == 0 {
		return c.Status(404).JSON(fiber.Map{
			"error": fmt.Sprintf("Library with ID %s not found", id),
		})	
	} else {
		return c.SendString("Library deleted successfully")
	}
}


// GET
func GetLibraries(c *fiber.Ctx) error {
	libraryCollection := database.GetCollection("libraries")

	cursor, err := libraryCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		return err
	}

	var libraries= []models.Library{}
	if err = cursor.All(context.TODO(), &libraries); err != nil {
		return err
	}

	//return c.Status(200).JSON(fiber.Map{"data": libraries})
	return c.JSON(libraries)
}


// GET
func GetLibrary(c *fiber.Ctx) error {
	libraryCollection := database.GetCollection("libraries")

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

	library := models.Library{}

	

	err = libraryCollection.FindOne(c.Context(), bson.M{"_id": objectId}).Decode(&library)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(404).JSON(fiber.Map{
				"error": fmt.Sprintf("Library with ID %s not found", id),
			})			
		} else {
			return c.Status(500).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
	}

	//return c.Status(200).JSON(fiber.Map{"data": library})
	return c.JSON(library)
}


type libraryDTO struct {
	Name    string   `json:"name" bson:"name"`
	Address string   `json:"address" bson:"address"`
	Empty   []string `json:"no_exists" bson:"books"`
}

// POST
func CreateLibrary(c *fiber.Ctx) error {
	nLibrary := new(libraryDTO)

	if err := c.BodyParser(nLibrary); err != nil {
		return err
	}

	nLibrary.Empty = make([]string, 0)

	libraryCollection := database.GetCollection("libraries")

	nDoc, err := libraryCollection.InsertOne(context.TODO(), nLibrary)
	if err != nil {
		return err
	}

	return c.Status(201).JSON(fiber.Map{"id": nDoc.InsertedID})
}
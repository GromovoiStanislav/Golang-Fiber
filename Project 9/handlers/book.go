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
	libraryCollection := database.GetCollection("libraries")

	// update one library
	_, err = libraryCollection.UpdateOne(context.TODO(), bson.M{"_id": objectId}, bson.D{{Key: "$push", Value: bson.M{"books": createData}}})
	if err != nil {
		return err
	}

	return c.Status(201).SendString("Book created successfully")
}


func GetBooks(c *fiber.Ctx) error {
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
	libraryCollection := database.GetCollection("libraries")

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
	return c.JSON(library.Books)
}


func GetBook2(c *fiber.Ctx) error {
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

	// get the bookId
	bookId := c.Params("bookId")
	if id == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "id is required",
		})
	}

	objectBookId, err := primitive.ObjectIDFromHex(bookId)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid id",
		})
	}


	// get the collection reference
	libraryCollection := database.GetCollection("libraries")

	library := models.Library{}

	filter := bson.D{
		{"_id", objectId},
		{"books", bson.D{{"$elemMatch", bson.D{{"_id", objectBookId}}}}},
	}

	err = libraryCollection.FindOne(c.Context(), filter).Decode(&library)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(404).JSON(fiber.Map{
				"error": fmt.Sprintf("Book with ID %s not found", bookId),
			})			
		} else {
			return c.Status(500).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
	}

	//return c.Status(200).JSON(fiber.Map{"data": library})
	return c.JSON(library.Books[0])
}


func GetBook(c *fiber.Ctx) error {
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

	// get the bookId
	bookId := c.Params("bookId")
	if id == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "id is required",
		})
	}

	objectBookId, err := primitive.ObjectIDFromHex(bookId)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid id",
		})
	}


	// get the collection reference
	libraryCollection := database.GetCollection("libraries")


	// Агрегация для поиска конкретной книги по её ID и ID библиотеки.
	pipeline := mongo.Pipeline{
		{{"$match", bson.D{{"_id", objectId}}}},
		{{"$unwind", "$books"}},
		{{"$match", bson.D{{"books._id", objectBookId}}}},
		{{"$project", bson.D{{"_id", "$books._id"}, {"title", "$books.title"}, {"author", "$books.author"} , {"year", "$books.year"}}}},
	}

	cur, err := libraryCollection.Aggregate(c.Context(), pipeline)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	defer cur.Close(c.Context())

	var result models.Book
	if cur.Next(context.Background()) {
		err := cur.Decode(&result)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.JSON(result)
	} else {
		return c.Status(404).JSON(fiber.Map{
			"error": fmt.Sprintf("Book with ID %s not found", bookId),
		})	
	}
}


func DeleteBook(c *fiber.Ctx) error {
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

	// get the bookId
	bookId := c.Params("bookId")
	if id == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "id is required",
		})
	}

	objectBookId, err := primitive.ObjectIDFromHex(bookId)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid id",
		})
	}


	// get the collection reference
	libraryCollection := database.GetCollection("libraries")

	filter := bson.D{
		{"_id", objectId},
		{"books", bson.D{{"$elemMatch", bson.D{{"_id", objectBookId}}}}},
	}

	update := bson.D{
		{"$pull", bson.D{
			{"books", bson.D{
				{"_id", objectBookId},
			}},
		}},
	}

	result, err := libraryCollection.UpdateOne(c.Context(), filter, update)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if result.ModifiedCount == 0 {
		return c.Status(404).JSON(fiber.Map{
			"error": fmt.Sprintf("Book with ID %s not found", bookId),
		})	
	} else {
		return c.SendString("Library deleted successfully")
	}
}
package book

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"main/database"
)

type Book struct {
	gorm.Model
	Title  string `json:"name"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}

func NewBook(c *fiber.Ctx) error {
	db := database.DBConn
	book := new(Book)

	if err := c.BodyParser(book); err != nil {
		c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"error": err.Error()})
		return err
	}

	result := db.Create(book)
	if result.Error != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
		return result.Error
	}

	return c.Status(fiber.StatusCreated).JSON(book)
}

func GetBooks(c *fiber.Ctx) error {
	db := database.DBConn
	var books []Book
	db.Find(&books)
	return c.JSON(books)
}

func GetBook(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	var book Book

	result := db.First(&book, id)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "No Book Found with ID"})
	}

	// Если использовать db.Find(&book, id) вместо db.First(&book, id)
	//if book.Title == "" {
	//	return c.Status(fiber.StatusNotFound).SendString("No Book Found with ID")
	//}

	return c.JSON(book)
}

func DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn

	var book Book
	result := db.First(&book, id)
	if result.Error != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "No Book Found with ID"})
	}

	// Если использовать db.Find(&book, id) вместо db.First(&book, id)
	//if book.Title == "" {
	//	return c.Status(fiber.StatusNotFound).SendString("No Book Found with ID")
	//}

	result = db.Delete(&book)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
	}

	return c.SendString("Book Successfully deleted")
}

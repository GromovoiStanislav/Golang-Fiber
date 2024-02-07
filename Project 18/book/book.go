package book

import (
	"github.com/gofiber/fiber/v2"

	"fiber-example/database"
	"fiber-example/models"
)


func GetBooks(c *fiber.Ctx) error {
	db := database.DBConn
	
	var books []models.Book
	db.Find(&books)
	
	return c.JSON(books)
}


func GetBook(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn
	
	var book models.Book
	db.Find(&book, id)
	if book.Title == "" {
        return c.Status(404).SendString("No Found The Book with ID") 
	}
	return c.JSON(book)
}


func NewBook(c *fiber.Ctx) error {
	db := database.DBConn

	book := new(models.Book)
	if err := c.BodyParser(book); err != nil {
		return c.Status(400).JSON(err)
	}
	
	db.Create(&book)

	return c.JSON(book)
}

func DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn

	var book models.Book
	db.Find(&book, id)
	if book.Title == "" {
        return c.Status(404).SendString("No Found The Book with ID") 
	}

	db.Delete(&book)
	return c.SendString("Book Successfully deleted")
}
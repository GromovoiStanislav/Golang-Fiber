package book

import (
	"github.com/gofiber/fiber/v2"

	"fiber-example/internal/database"
	"fiber-example/internal/models"
)


type Message struct {
	Msg string 
}


func GetBooks(c *fiber.Ctx) error {
	db := database.DBConn
	
	var books []models.Book
	db.Find(&books)
	
	return c.JSON(books)
}


func GetBook(c *fiber.Ctx) error {
	db := database.DBConn

	ISIN := c.Params("isin")

	var book models.Book
	// result := db.Where("ISIN = ?", ISIN).Find(&book)
	result := db.Find(&book, "ISIN = ?", ISIN)
	if result.RowsAffected == 0 {
        return c.Status(404).JSON(Message{Msg: "No Found The Book with ISIN"}) 
    }

	return c.Status(fiber.StatusOK).JSON(book)
}


func NewBook(c *fiber.Ctx) error {
	db := database.DBConn

	book := new(models.Book)
	if err := c.BodyParser(book); err != nil {
		return c.Status(400).JSON(err)
	}
	
	result := db.Create(&book)
	if result.Error != nil {
		return c.Status(fiber.StatusBadRequest).JSON(Message{Msg: "Bad Request"}) 
	}
	// if result.Error != nil {
	// 	c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": result.Error.Error()})
	// 	return result.Error
	// }

	return c.Status(fiber.StatusCreated).JSON(book)
}

func DeleteBook(c *fiber.Ctx) error {
	ISIN := c.Params("isin")
	db := database.DBConn

	var book models.Book
	
	// result := db.Find(&book, "ISIN = ?", ISIN)
	// if result.RowsAffected == 0 {
    //     return c.Status(404).JSON(Message{Msg: "No Found The Book with ISIN"}) 
    // }
	// db.Delete(&book)

	result := db.Where("ISIN = ?", ISIN).Delete(&book)
	//result := db.Delete(&book, "ISIN = ?", ISIN)
	if result.RowsAffected == 0 {
        return c.Status(fiber.StatusNotFound).JSON(Message{Msg: "No Found The Book with ISIN"}) 
    }

	return c.Status(fiber.StatusOK).JSON(Message{Msg: "Book Successfully deleted"}) 
}
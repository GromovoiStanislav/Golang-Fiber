package main

import (
	"os"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/go-playground/validator/v10"
	_ "github.com/joho/godotenv/autoload"

	"fiber-example/bookmark"
	"fiber-example/database"
	"fiber-example/utils"
)



func main() {
	app := fiber.New()
	dbErr := database.InitDB()

	utils.Validator = validator.New(validator.WithRequiredStructEnabled())

	if dbErr != nil {
		panic(dbErr)
	}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/api/bookmarks", bookmark.GetAllBookmarks)
	app.Get("/api/bookmarks/:id", bookmark.GetBookmark)
	app.Post("/api/bookmarks", bookmark.ValidateCreateBookmark, bookmark.NewBookmark)
	app.Put("/api/bookmarks/:id", bookmark.ValidateCreateBookmark,bookmark.UpdateBookmarkPut)
	app.Patch("/api/bookmarks/:id", bookmark.ValidateUpdateBookmark,bookmark.UpdateBookmarkPatch)
	app.Delete("/api/bookmarks/:id", bookmark.DeleteBookmark)

	PORT := os.Getenv("PORT")
	app.Listen(fmt.Sprintf(":%s", PORT))
}
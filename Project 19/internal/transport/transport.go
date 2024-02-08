package transport

import (
	"github.com/gofiber/fiber/v2"

	"fiber-example/internal/book"
)

func helloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello, World!")
}

func setupRoutes(app *fiber.App) {
	app.Get("/", helloWorld)
	app.Get("/api/v1/book", book.GetBooks)
	app.Get("/api/v1/book/:isin", book.GetBook)
	app.Post("/api/v1/book", book.NewBook)
	app.Delete("/api/v1/book/:isin", book.DeleteBook)
}

func Setup() *fiber.App {
	app := fiber.New()
	setupRoutes(app)
	return app
}
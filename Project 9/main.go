package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	"fiber-example/database"
	"fiber-example/handlers"
)

func main() {
	// init app
	err := initApp()
	if err != nil {
		panic(err)
	}

	// defer close database
	defer database.CloseMongoDB()

	app := generateApp()

	// get the port from the env
	port := os.Getenv("PORT")
	app.Listen(":" + port)
}

func generateApp() *fiber.App {
	app := fiber.New()

	// create health check route
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	// create the library group and routes
	libGroup := app.Group("/library")
	libGroup.Get("/", handlers.GetLibraries)
	libGroup.Get("/:id", handlers.GetLibrary)
	libGroup.Post("/", handlers.CreateLibrary)
	libGroup.Delete("/:id", handlers.DeleteLibrary)
	
	libGroup.Post("/:id/books", handlers.CreateBook)
	libGroup.Get("/:id/books", handlers.GetBooks)
	libGroup.Get("/:id/books/:bookId", handlers.GetBook)
	libGroup.Delete("/:id/books/:bookId", handlers.DeleteBook)
	
	return app
}

func initApp() error {
	// setup env
	err := loadENV()
	if err != nil {
		return err
	}

	// setup database
	err = database.StartMongoDB()
	if err != nil {
		return err
	}

	return nil
}

func loadENV() error {
	// check if prod
	prod := os.Getenv("PROD")
	if prod != "true" {
		err := godotenv.Load()
		if err != nil {
			return err
		}
	}
	return nil
}
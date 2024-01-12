package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"

	"fiber-example/internal/handlers"
)

func healthcheck(c *fiber.Ctx) error {
	return c.SendString("OK")
}

func main() {

	app := fiber.New()

	app.Use("/api", func(c *fiber.Ctx) error {
		fmt.Println("Hello from middleware")
		return c.Next()
	})

	app.Get("/healthcheck", healthcheck)

	app.Post("/api/products", handlers.CreateProduct)
	app.Get("/api/products", handlers.GetAllProducts)

	log.Fatal(app.Listen(":3000"))

}
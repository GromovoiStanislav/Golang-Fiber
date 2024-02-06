package main

import (
	"os"
	"fmt"
	"log"


	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"

	"fiber-example/config"
	"fiber-example/handlers"
)


func main() {
    config.Connect()

    app := fiber.New()

    app.Get("/dogs", handlers.GetDogs)
    app.Get("/dogs/:id", handlers.GetDog)
    app.Post("/dogs", handlers.AddDog)
    app.Put("/dogs/:id", handlers.UpdateDog)
    app.Delete("/dogs/:id", handlers.RemoveDog)

	PORT := os.Getenv("PORT")
    log.Fatal(app.Listen(fmt.Sprintf(":%s", PORT)))
}
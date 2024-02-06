package main

import (
	"os"
	"context"

	_ "github.com/joho/godotenv/autoload"
	"github.com/gofiber/fiber/v2"
	
    "fiber-example/config"
    "fiber-example/handlers"
)


type Person struct {
    Name  string
    Phone string
}

func main() {
    app := fiber.New()
    if err := config.ConnectDatabase(); err != nil {
        panic(err)
    }
    defer config.Client.Disconnect(context.Background())
    app.Get("/dogs", handlers.GetDogs)
    app.Get("/dogs/:id", handlers.GetDog)
    app.Post("/dogs", handlers.AddDog)
    app.Put("/dogs/:id", handlers.UpdateDog)
    app.Delete("/dogs/:id", handlers.RemoveDog)

    // get the port from the env
	port := os.Getenv("PORT")
	app.Listen(":" + port)
}
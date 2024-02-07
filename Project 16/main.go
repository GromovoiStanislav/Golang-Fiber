package main

import (
	"os"
	"fmt"


	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/joho/godotenv/autoload"

	"fiber-example/routes"
	"fiber-example/database"
)



func main() {
	database.Connect()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.Setup(app)

	PORT := os.Getenv("PORT")
	app.Listen(fmt.Sprintf(":%s", PORT))
}
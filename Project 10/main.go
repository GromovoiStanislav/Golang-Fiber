package main

import (
	"os"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"

	"fiber-example/routes"
	"fiber-example/utils"
)

func main() {
	PORT := os.Getenv("PORT")
	utils.InitDB()

	app := fiber.New()
	todos := fiber.New()

	app.Mount("/api/todos", todos)
	routes.Todos(todos)

	utils.Validator = validator.New(validator.WithRequiredStructEnabled())

	app.Listen(fmt.Sprintf(":%s", PORT))
}
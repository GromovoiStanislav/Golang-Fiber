package main

import (
	"os"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/go-playground/validator/v10"
	_ "github.com/joho/godotenv/autoload"

	"fiber-example/config"
	"fiber-example/routes"
	"fiber-example/utils"
)


func main() {
	config.ManageMigrations()

	utils.Validator = validator.New(validator.WithRequiredStructEnabled())

	app := fiber.New()

	routes.InitRouters(app)

	PORT := os.Getenv("PORT")
	err := app.Listen(fmt.Sprintf(":%s", PORT))
	if err != nil {
		log.Fatalln("Router HTTP error: ", err)
	}
}
package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"fiber-example/src/config"
	"fiber-example/src/routes"
)

func main() {
	// Initializations
	app := fiber.New()
	
	// Static files
	app.Static("/", "./public")

	// Routes
	routes.IndexRoutes(app)
	routes.UserRoute(app)

	cfg := config.GetConfig()
	app.Listen(fmt.Sprintf(":%s", cfg.PORT))
}
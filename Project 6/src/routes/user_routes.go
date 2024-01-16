package routes

import (
	"github.com/gofiber/fiber/v2"

	"fiber-example/src/controllers"
)

func UserRoute(app *fiber.App) {
	app.Post("/users", controllers.CreateUser)
	app.Get("/users/:userId", controllers.GetUser)
	app.Delete("/users/:userId", controllers.DeleteUser)
	app.Get("/users", controllers.GetUsers)
	app.Patch("/users/:userId", controllers.UpdateUser)
}
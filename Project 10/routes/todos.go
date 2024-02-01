package routes

import (
	"github.com/gofiber/fiber/v2"

	"fiber-example/handlers"
)

func Todos(app *fiber.App) {
	app.Get("/:id", handlers.GetTodoHandler)
	app.Get("/", handlers.GetTodosHandler)
	app.Post("/", handlers.ValidateCreateTodo, handlers.CreateTodoHandler)
}
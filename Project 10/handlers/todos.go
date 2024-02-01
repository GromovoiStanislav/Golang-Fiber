package handlers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"

	"fiber-example/handlers/dtos"
	"fiber-example/models"
	"fiber-example/utils"

)

func GetTodosHandler(ctx *fiber.Ctx) error {
	var todos []models.Todo

	if result := utils.DB.Find(&todos); result.Error != nil {
		fmt.Println(result.Error)
		return ctx.SendStatus(500)
	}

	return ctx.Status(200).JSON(todos)
}

func GetTodoHandler(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	
	var todo models.Todo


	// if result := utils.DB.First(&todo,id); result.Error != nil {
	// 	return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Not Found todo by ID"})
	// }

	// OR:

	if result := utils.DB.Find(&todo,id); result.Error != nil {
		fmt.Println(result.Error)
		return ctx.SendStatus(500)
	}
	if todo.Title == "" {
		return ctx.Status(fiber.StatusNotFound).SendString("Not Found todo by ID")
	}

	return ctx.Status(200).JSON(todo)
}

func ValidateCreateTodo(ctx *fiber.Ctx) error {
	createTodo := &dtos.CreateTodoDto{}
	if err := ctx.BodyParser(createTodo); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"msg": "Error"})
	}
	if err := utils.Validator.Struct(createTodo); err != nil {
		fmt.Println(err)
		return ctx.Status(400).JSON(fiber.Map{"msg": "Bad Request"})
	}
	return ctx.Next()
}

func CreateTodoHandler(ctx *fiber.Ctx) error {
	newTodo := new(models.Todo)
	if err := ctx.BodyParser(&newTodo); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"msg": "Error"})
	}
	if result := utils.DB.Create(&newTodo); result.Error != nil {
		fmt.Println(result.Error)
		return ctx.SendStatus(500)
	}
	return ctx.Status(201).JSON(newTodo)
}
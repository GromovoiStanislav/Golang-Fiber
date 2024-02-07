package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"

	"fiber-example/utils"
	"fiber-example/models"
)


func ValidateBlogDto(ctx *fiber.Ctx) error {
	//blog := &models.BlogDto{}
	blog := new(models.BlogDto)

	if err := ctx.BodyParser(blog); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"msg": "Error"})
	}

	if err := utils.Validator.Struct(blog); err != nil {
		fmt.Println(err)
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"msg": "Bad Request"})
	}
	return ctx.Next()
}


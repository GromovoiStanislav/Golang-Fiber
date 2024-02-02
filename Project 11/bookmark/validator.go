package bookmark

import (
	"fmt"
	"github.com/gofiber/fiber/v2"

	"fiber-example/utils"

)

type CreateBookmarkDto struct {
	Name       string `json:"name" validate:"required"`
	URL string `json:"url" validate:"required"`
}


type UpdateBookmarkDto struct {
	Name       string `json:"name" validate:""`
	URL string `json:"url" validate:""`
}


func ValidateCreateBookmark(ctx *fiber.Ctx) error {
	createBookmark := &CreateBookmarkDto{}
	if err := ctx.BodyParser(createBookmark); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"msg": "Error"})
	}
	if err := utils.Validator.Struct(createBookmark); err != nil {
		fmt.Println(err)
		return ctx.Status(400).JSON(fiber.Map{"msg": "Bad Request"})
	}
	return ctx.Next()
}

func ValidateUpdateBookmark(ctx *fiber.Ctx) error {
	updateBookmark := &UpdateBookmarkDto{}
	if err := ctx.BodyParser(updateBookmark); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"msg": "Error"})
	}
	if err := utils.Validator.Struct(updateBookmark); err != nil {
		fmt.Println(err)
		return ctx.Status(400).JSON(fiber.Map{"msg": "Bad Request"})
	}
	return ctx.Next()
}
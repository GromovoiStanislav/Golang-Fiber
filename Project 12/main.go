package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/go-playground/validator/v10"
)

type Dog struct {
    Name      string `json:"name" validate:"required,min=3,max=12"`
    Age       int    `json:"age" validate:"required,numeric"`
    IsGoodBoy bool   `json:"isGoodBoy" validate:"required"`
}

type IError struct {
    Field string
    Tag   string
    Value string
}

var Validator = validator.New()

func main() {
    app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Thank god it works 🙏")
    })

	app.Post("/", ValidateAddDog, func(c *fiber.Ctx) error {
        body := new(Dog)
        c.BodyParser(&body)
        return c.Status(fiber.StatusOK).JSON(body)
    })

    app.Listen(":3000")
}

func ValidateAddDog(c *fiber.Ctx) error {
    var errors []*IError
    body := new(Dog)
    c.BodyParser(&body)

    err := Validator.Struct(body)
    if err != nil {
        for _, err := range err.(validator.ValidationErrors) {
            var el IError
            el.Field = err.Field()
            el.Tag = err.Tag()
            el.Value = err.Param()
            errors = append(errors, &el)
        }
        return c.Status(fiber.StatusBadRequest).JSON(errors)
    }
    return c.Next()
}
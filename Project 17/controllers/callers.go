package controllers

import (
	"github.com/gofiber/fiber/v2"

	"fiber-example/config"
	"fiber-example/models"
)


func GetBlogs(ctx *fiber.Ctx) error {
	var blogs []models.Blogs

	config.DB.Find(&blogs)

	return ctx.Status(fiber.StatusOK).JSON(&blogs)	
}



func GetBlog(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	var blog models.Blogs

	result := config.DB.Find(&blog, id)
	if result.RowsAffected == 0 {
		return ctx.SendStatus(fiber.StatusNotFound)
	}

	return ctx.Status(fiber.StatusOK).JSON(&blog)
}


func CreateBlog(ctx *fiber.Ctx) error {
	blog := new(models.Blogs)
	
	if err := ctx.BodyParser(blog); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}

	config.DB.Create(&blog)

	return ctx.Status(fiber.StatusCreated).JSON(blog)
}


func DeleteBlog(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	var blog models.Blogs

	search := config.DB.Delete(&blog, id)
	if search.RowsAffected == 0 {
		return ctx.Status(fiber.StatusNotFound).JSON(models.Message{
			Status:      "Error",
			Description: "404: Not Found",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(models.Message{
		Status:      "Success",
		Description: "The post was successfully deleted",
	})
}


func UpdateBlog(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	blog := new(models.Blogs)

	if err := ctx.BodyParser(blog); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(models.Message{
			Status:      "Error",
			Description: "There was an error parsing your request",
		})
	}

	result := config.DB.Where("id = ?", id).Updates(&blog)
	if result.RowsAffected == 0 {
			return ctx.SendStatus(fiber.StatusNotFound)
	}

	updatedBlog := new(models.Blogs)
	config.DB.Find(&updatedBlog, id)

	return ctx.Status(fiber.StatusOK).JSON(updatedBlog)
}
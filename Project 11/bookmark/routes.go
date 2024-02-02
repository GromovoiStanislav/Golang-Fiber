package bookmark

import (
	"github.com/gofiber/fiber/v2"

	"fiber-example/database"
)

func GetAllBookmarks(c *fiber.Ctx) error {
	result, err := database.GetAllBookmarks()

	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"message": "Internal Server Error",
			"data":    nil,
		})
	}

	return c.Status(200).JSON(&fiber.Map{
		"message": "Success",
		"data":    result,
	})
}

func GetBookmark(c *fiber.Ctx) error {
	result, err := database.GetBookmark(c.Params("id"))

	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"message": "Internal Server Error",
			"data":    nil,
		})
	}

	if result.ID == 0 {
		return c.Status(404).JSON(&fiber.Map{
			"message": "Bookmark not found",
			"data":    nil,
		})
	}

	return c.Status(200).JSON(&fiber.Map{
		"message": "Success",
		"data":    result,
	})
}

func NewBookmark(c *fiber.Ctx) error {
	bookmark := database.Bookmark{}


	if err := c.BodyParser(&bookmark); err != nil {
		c.Status(400).JSON(&fiber.Map{
			"message": "Invalid Input Data",
			"data":    nil,
		})
		return err
	}


	result, err := database.CreateBookmark(bookmark.Name, bookmark.URL)

	if err != nil {
		c.Status(500).JSON(&fiber.Map{
			"message": "Internal Server Error",
			"data":    nil,
		})
		return err
	}

	c.Status(201).JSON(&fiber.Map{
		"message": "Bookmark created successfully",
		"data":    result,
	})

	return nil
}

func UpdateBookmarkPut(c *fiber.Ctx) error {
	var bookmark database.Bookmark

	if err := c.BodyParser(&bookmark); err != nil {
		c.Status(400).JSON(&fiber.Map{
			"message": "Invalid Input Data",
			"data":    nil,
		})
		return err
	}

	result, err := database.UpdateBookmarkPut(c.Params("id"), bookmark.Name, bookmark.URL)

	if err != nil {

		if err.Error()=="record not found" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Bookmark not found",
				"data":    nil,
			})
		}

		c.Status(500).JSON(&fiber.Map{
			"message": "Internal Server Error",
			"data":    nil,
		})
		return err
	}

	return c.Status(200).JSON(&fiber.Map{
		"message": "Bookmark updated successfully",
		"data":    result,
	})
}

func UpdateBookmarkPatch(c *fiber.Ctx) error {
	var bookmark database.Bookmark

	if err := c.BodyParser(&bookmark); err != nil {
		c.Status(400).JSON(&fiber.Map{
			"message": "Invalid Input Data",
			"data":    nil,
		})
		return err
	}

	result, err := database.UpdateBookmarkPatch(c.Params("id"), bookmark.Name, bookmark.URL)

	if err != nil {

		if err.Error()=="record not found" {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": "Bookmark not found",
				"data":    nil,
			})
		}

		c.Status(500).JSON(&fiber.Map{
			"message": "Internal Server Error",
			"data":    nil,
		})
		return err
	}

	return c.Status(200).JSON(&fiber.Map{
		"message": "Bookmark updated successfully",
		"data":    result,
	})
}

func DeleteBookmark(c *fiber.Ctx) error {
	result, err := database.DeleteBookmark(c.Params("id"))

	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"message": "Internal Server Error",
			"data":    nil,
		})
	}

	return c.Status(200).JSON(&fiber.Map{
		"message": "Bookmark deleted successfully",
		"data":    result,
	})
}
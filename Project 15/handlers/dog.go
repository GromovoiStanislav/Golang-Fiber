package handlers

import (
    "github.com/gofiber/fiber/v2"

    "fiber-example/config"
    "fiber-example/entities"
)


func GetDogs(c *fiber.Ctx) error {
    var dogs []entities.Dog

    config.Database.Find(&dogs)
    return c.Status(200).JSON(dogs)
}


func GetDog(c *fiber.Ctx) error {
    id := c.Params("id")
    var dog entities.Dog

    result := config.Database.Find(&dog, id)

    if result.RowsAffected == 0 {
        return c.SendStatus(404)
    }

    return c.Status(200).JSON(&dog)
}


func AddDog(c *fiber.Ctx) error {
    dog := new(entities.Dog)

    if err := c.BodyParser(dog); err != nil {
        return c.Status(503).SendString(err.Error())
    }

    config.Database.Create(&dog)
    return c.Status(201).JSON(dog)
}


func UpdateDog(c *fiber.Ctx) error {
    dog := new(entities.Dog)
    id := c.Params("id")

    if err := c.BodyParser(dog); err != nil {
        return c.Status(503).SendString(err.Error())
    }

    result := config.Database.Where("id = ?", id).Updates(&dog)
    if result.RowsAffected == 0 {
        return c.SendStatus(404)
    }
    

    // Загрузка обновленного объекта из базы данных
    updatedDog := new(entities.Dog)
    config.Database.First(&updatedDog, id)

    return c.Status(200).JSON(updatedDog)
}

func RemoveDog(c *fiber.Ctx) error {
    id := c.Params("id")
    var dog entities.Dog

    result := config.Database.Delete(&dog, id)

    if result.RowsAffected == 0 {
        return c.SendStatus(404)
    }

    return c.SendStatus(200)
}
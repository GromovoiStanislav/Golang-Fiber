package handlers

import (
    "github.com/gofiber/fiber/v2"
    "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
    "go.mongodb.org/mongo-driver/bson"

	"fiber-example/config"
    "fiber-example/entities"
)


func GetDogs(c *fiber.Ctx) error {
    query := bson.D{}
    data, err := config.Collections.Dogs.Find(c.Context(), query)
    if err != nil {
        return c.Status(500).SendString(err.Error())
    }
    var allDogs []entities.IDogs = make([]entities.IDogs, 0)
    if err := data.All(c.Context(), &allDogs); err != nil {
        return c.Status(500).JSON(err)
    }
    return c.Status(200).JSON(&allDogs)
}


func GetDog(c *fiber.Ctx) error {
    dogId, err := primitive.ObjectIDFromHex(c.Params("id"))
    if err != nil {
        return c.Status(500).JSON(err)
    }
    var singleDog bson.M
    query := bson.D{{"_id", dogId}}
    if err := config.Collections.Dogs.FindOne(c.Context(), query).Decode(&singleDog); err != nil {
        return c.Status(500).JSON(err)
    }

    return c.Status(200).JSON(singleDog)
}


func AddDog(c *fiber.Ctx) error {
    dog := new(entities.IDogs)
    c.BodyParser(dog)
    response, err := config.Collections.Dogs.InsertOne(c.Context(), dog)
    if err != nil {
        return c.Status(500).JSON(err)
    }
    var insertedDog bson.M
    query := bson.D{{"_id", response.InsertedID}}
    if err := config.Collections.Dogs.FindOne(c.Context(), query).Decode(&insertedDog); err != nil {
        return c.Status(500).JSON(err)
    }
    return c.Status(200).JSON(insertedDog)
}


func UpdateDog(c *fiber.Ctx) error {
    dogId, err := primitive.ObjectIDFromHex(
        c.Params("id"),
    )
    if err != nil {
        return c.Status(500).JSON(err)
    }
    dog := new(entities.IDogs)
    if err := c.BodyParser(dog); err != nil {
        return c.Status(500).JSON(err)
    }
    query := bson.D{{"_id", dogId}}
    body := bson.D{
        {Key: "$set",
            Value: bson.D{
                {"name", dog.Name},
                {"breed", dog.Breed},
                {"age", dog.Age},
                {"isGoodBoy", dog.IsGoodBoy},
            },
        },
    }
    
    // if _, err := config.Collections.Dogs.UpdateOne(c.Context(), &query, &body); err != nil {
    //     return c.Status(500).JSON(err)
    // }
    //var updatedDog bson.M
    // if err := config.Collections.Dogs.FindOne(c.Context(), &query).Decode(&updatedDog); err != nil {
    //     return c.Status(500).JSON(err)
    // }

    var updatedDog bson.M
	updateOptions := options.FindOneAndUpdate().SetReturnDocument(options.After)
    if err := config.Collections.Dogs.FindOneAndUpdate(c.Context(), &query, &body, updateOptions).Decode(&updatedDog); err != nil {
        return c.Status(500).JSON(err)
    }

    return c.Status(200).JSON(updatedDog)
}


func RemoveDog(c *fiber.Ctx) error {
    dogId, err := primitive.ObjectIDFromHex(
        c.Params("id"),
    )
    if err != nil {
        return c.Status(500).JSON(err)
    }

    query := bson.D{{"_id", dogId}}
    var dog bson.M
    if err := config.Collections.Dogs.FindOneAndDelete(c.Context(), &query).Decode(&dog); err != nil {
        return c.Status(500).JSON(err)
    }
    return c.Status(200).JSON(dog)
}
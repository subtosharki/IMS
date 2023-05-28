package users

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"inventory-api/src/utils/structs"
)

func GetByAPIKey(c *fiber.Ctx) error {
	apikey := c.Params("apikey")
	if apikey == "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "User not found"})
	}

	user := new(structs.User)
	mongodb := c.Locals("mongo").(*mongo.Database)
	collection := mongodb.Collection("users")
	err := collection.FindOne(context.Background(), bson.M{"apikey": apikey}).Decode(&user)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "User not found"})
	}
	user.Clean()
	return c.JSON(user)
}

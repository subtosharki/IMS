package clones

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"inventory-api/src/utils/structs"
)

func Delete(c *fiber.Ctx) error {
	user := c.Locals("user").(*structs.User)
	if user.Role != "admin" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Unauthorized"})
	}

	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Missing clone id"})
	}

	mongodb := c.Locals("mongo").(*mongo.Database)
	collection := mongodb.Collection("clones")

	result, err := collection.DeleteOne(context.Background(), bson.M{"cloneId": id})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error deleting clone"})
	}

	if result.DeletedCount == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Clone not found"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Clone deleted successfully"})
}

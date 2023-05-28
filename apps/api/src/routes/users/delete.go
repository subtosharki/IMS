package users

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
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"message": "You do not have permission to delete users"})
	}

	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Missing user id"})
	}

	if user.UserId == id {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"message": "You cannot delete yourself"})
	}

	mongodb := c.Locals("mongo").(*mongo.Database)
	collection := mongodb.Collection("users")

	result, err := collection.DeleteOne(context.Background(), bson.M{"userId": id})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error deleting user"})
	}

	if result.DeletedCount == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "User not found"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "User deleted successfully"})
}

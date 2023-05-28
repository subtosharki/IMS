package users

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"inventory-api/src/utils/structs"
)

func ToggleAdmin(c *fiber.Ctx) error {
	performingUser := c.Locals("user").(*structs.User)
	if performingUser.Role != "admin" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"message": "You do not have permission to toggle admin"})
	}

	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Missing user id"})
	}

	mongodb := c.Locals("mongo").(*mongo.Database)
	collection := mongodb.Collection("users")

	user := new(structs.User)

	err := collection.FindOne(context.Background(), bson.M{"userId": id}).Decode(&user)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "User not found"})
	}

	role := "sales"
	if user.Role == "sales" {
		role = "admin"
	}

	_, err = collection.UpdateOne(context.Background(), bson.M{"userId": id}, bson.M{"$set": bson.M{"role": role}})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error updating user"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "User updated successfully"})
}

package customers

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Delete(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Missing customer id"})
	}

	mongodb := c.Locals("mongo").(*mongo.Database)
	collection := mongodb.Collection("customers")

	result, err := collection.DeleteOne(context.Background(), bson.M{"customerId": id})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error deleting customer"})
	}

	if result.DeletedCount == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Customer not found"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Customer deleted successfully"})
}

package clones

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"inventory-api/src/utils/structs"
	"strconv"
)

func SetQuantity(c *fiber.Ctx) error {
	user := c.Locals("user").(*structs.User)
	if user.Role != "admin" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Unauthorized"})
	}

	cloneId := c.Params("id")
	if cloneId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Missing clone id"})
	}
	quantity, err := strconv.ParseFloat(c.Params("quantity"), 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid quantity"})
	}
	if quantity < 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid quantity"})
	}

	mongodb := c.Locals("mongo").(*mongo.Database)
	collection := mongodb.Collection("clones")

	clone := new(structs.Clone)
	err = collection.FindOneAndUpdate(context.Background(), bson.M{"cloneId": cloneId}, bson.M{"$set": bson.M{"quantity": quantity}}).Decode(&clone)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error updating clone"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Clone updated successfully"})
}

package middleware

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"inventory-api/src/utils/structs"
)

func APIKeyAuth(c *fiber.Ctx) error {
	req := new(structs.APIKeyRequest)
	err := c.ReqHeaderParser(req)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid API key"})
	}

	if req.APIKey == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid API key"})
	}

	mongodb := c.Locals("mongo").(*mongo.Database)
	collection := mongodb.Collection("users")
	user := new(structs.User)
	err = collection.FindOne(context.Background(), bson.M{"apikey": req.APIKey}).Decode(&user)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid API key"})
	}

	c.Locals("user", user)
	return c.Next()

}

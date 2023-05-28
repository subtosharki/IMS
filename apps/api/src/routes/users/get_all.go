package users

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"inventory-api/src/utils/structs"
)

func GetAll(c *fiber.Ctx) error {
	mongodb := c.Locals("mongo").(*mongo.Database)
	collection := mongodb.Collection("users")
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		panic(err)
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			panic(err)
		}
	}(cursor, context.Background())

	var users []structs.User
	for cursor.Next(context.Background()) {
		var user structs.User
		err := cursor.Decode(&user)
		if err != nil {
			panic(err)
		}
		users = append(users, user)
	}
	return c.Status(fiber.StatusOK).JSON(users)
}

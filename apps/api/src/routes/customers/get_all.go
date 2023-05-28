package customers

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"inventory-api/src/utils/structs"
)

func GetAll(c *fiber.Ctx) error {
	mongodb := c.Locals("mongo").(*mongo.Database)
	collection := mongodb.Collection("customers")
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

	var customers []structs.Customer
	for cursor.Next(context.Background()) {
		var customer structs.Customer
		err := cursor.Decode(&customer)
		if err != nil {
			panic(err)
		}
		customers = append(customers, customer)
	}
	return c.Status(fiber.StatusOK).JSON(customers)
}

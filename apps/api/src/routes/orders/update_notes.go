package orders

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"inventory-api/src/utils/structs"
)

func UpdateNotes(c *fiber.Ctx) error {
	orderId := c.Params("id")
	if orderId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Missing order id"})
	}

	bodyRequest := new(structs.UpdateNotesRequest)
	err := c.BodyParser(bodyRequest)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid request"})
	}

	errors := validateUpdateNotesBody(*bodyRequest)
	if len(errors) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid JSON fields", "errors": errors})
	}

	mongodb := c.Locals("mongo").(*mongo.Database)
	collection := mongodb.Collection("orders")
	result, err := collection.UpdateOne(context.Background(), bson.M{"orderNumber": orderId}, bson.M{"$set": bson.M{"notes": bodyRequest.Notes}})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error saving notes"})
	}

	if result.ModifiedCount == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Notes are identical to the current notes"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Notes saved successfully"})
}

func validateUpdateNotesBody(body structs.UpdateNotesRequest) []*structs.ValidationError {
	validate := validator.New()
	var errors []*structs.ValidationError
	err := validate.Struct(body)
	if err != nil {
		for _, validationErr := range err.(validator.ValidationErrors) {
			var element structs.ValidationError
			element.FailedField = validationErr.StructNamespace()
			element.Tag = validationErr.Tag()
			element.Value = validationErr.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}

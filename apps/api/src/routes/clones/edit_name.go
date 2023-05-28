package clones

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"inventory-api/src/utils/structs"
)

func EditName(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Missing clone id"})
	}

	bodyRequest := new(structs.EditNameRequest)
	err := c.BodyParser(bodyRequest)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid request"})
	}

	errors := validateEditCloneNameBody(*bodyRequest)
	if len(errors) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid JSON fields", "errors": errors})
	}

	if bodyRequest.Name == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid name"})
	}

	mongodb := c.Locals("mongo").(*mongo.Database)
	collection := mongodb.Collection("clones")

	clone := new(structs.Clone)
	err = collection.FindOneAndUpdate(context.Background(), bson.M{"cloneId": id}, bson.M{"$set": bson.M{"name": bodyRequest.Name}}).Decode(&clone)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error updating clone name"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Clone name updated successfully"})
}

func validateEditCloneNameBody(body structs.EditNameRequest) []*structs.ValidationError {
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

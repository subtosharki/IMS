package clones

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"inventory-api/src/utils/structs"
)

func Create(c *fiber.Ctx) error {
	user := c.Locals("user").(*structs.User)
	if user.Role != "admin" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Unauthorized"})
	}

	bodyRequest := new(structs.CreateCloneRequest)

	err := c.BodyParser(bodyRequest)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Error parsing body"})
	}

	errors := validateCreateClonesBody(*bodyRequest)
	if len(errors) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid JSON fields", "errors": errors})
	}

	mongodb := c.Locals("mongo").(*mongo.Database)
	collection := mongodb.Collection("clones")

	result, err := collection.InsertOne(context.Background(), bson.M{"name": bodyRequest.Name, "quantity": bodyRequest.Quantity, "cloneId": uuid.New().String(), "date": bodyRequest.Date})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error creating clone"})
	}

	if result.InsertedID == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error creating clone"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Clone created successfully"})
}

func validateCreateClonesBody(body structs.CreateCloneRequest) []*structs.ValidationError {
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

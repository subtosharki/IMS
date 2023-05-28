package users

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"inventory-api/src/utils/structs"
)

func Create(c *fiber.Ctx) error {
	bodyRequest := new(structs.CreateUserRequest)

	err := c.BodyParser(bodyRequest)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Error parsing body"})
	}

	errors := validateCreateUserBody(*bodyRequest)
	if len(errors) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid JSON fields", "errors": errors})
	}

	user := new(structs.User)
	user.GenerateUserID()
	user.GenerateAPIKey()

	mongodb := c.Locals("mongo").(*mongo.Database)
	collection := mongodb.Collection("users")

	result, err := collection.InsertOne(context.Background(), bson.M{"email": bodyRequest.Email, "password": bodyRequest.Password, "admin": bodyRequest.Admin, "apikey": user.APIKey})

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error creating user"})
	}

	if result.InsertedID == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error creating user"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "User created successfully"})
}

func validateCreateUserBody(body structs.CreateUserRequest) []*structs.ValidationError {
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

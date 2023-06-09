package auth

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"inventory-api/src/utils/structs"
)

func Login(c *fiber.Ctx) error {
	bodyRequest := new(structs.LoginRequest)
	err := c.BodyParser(bodyRequest)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid JSON fields", "error": err.Error()})
	}
	errors := validateLoginBody(*bodyRequest)
	if len(errors) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid JSON fields", "errors": errors})
	}

	user := new(structs.User)
	mongodb := c.Locals("mongo").(*mongo.Database)
	collection := mongodb.Collection("users")
	err = collection.FindOne(context.Background(), bson.M{
		"email":    bodyRequest.Email,
		"password": bodyRequest.Password,
	}).Decode(&user)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "Invalid credentials", "error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"apikey": user.APIKey})
}

func validateLoginBody(body structs.LoginRequest) []*structs.ValidationError {
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

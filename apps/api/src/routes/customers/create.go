package customers

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"inventory-api/src/utils/structs"
)

func Create(c *fiber.Ctx) error {
	bodyRequest := new(structs.CreateCustomerRequest)
	err := c.BodyParser(bodyRequest)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid JSON fields", "error": err.Error()})
	}

	errors := validateCreateCustomerBody(*bodyRequest)
	if len(errors) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid JSON fields", "errors": errors})
	}

	mongodb := c.Locals("mongo").(*mongo.Database)
	collection := mongodb.Collection("customers")

	customer := new(structs.Customer)
	customer.GenerateCustomerId()

	result, err := collection.InsertOne(context.Background(), bson.M{
		"customerId":     customer.CustomerId,
		"name":           bodyRequest.Name,
		"email":          bodyRequest.Email,
		"subscriptionId": bodyRequest.SubscriptionId,
		"phone":          bodyRequest.Phone,
		"address":        bodyRequest.Address,
		"notes":          ""})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error when Creating Customer", "error": err.Error()})
	}

	if result.InsertedID == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error when Creating Customer"})

	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Customer created successfully"})
}

func validateCreateCustomerBody(body structs.CreateCustomerRequest) []*structs.ValidationError {
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

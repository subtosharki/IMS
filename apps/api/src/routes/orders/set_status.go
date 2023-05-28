package orders

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"inventory-api/src/utils/structs"
)

func SetStatus(c *fiber.Ctx) error {
	orderId := c.Params("id")
	if orderId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Missing order id"})
	}

	bodyRequest := new(structs.SetStatusRequest)
	err := c.BodyParser(bodyRequest)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid request"})
	}

	errors := validateSetStatusBody(*bodyRequest)
	if len(errors) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid JSON fields", "errors": errors})
	}

	if bodyRequest.Status != "Voided" && bodyRequest.Status != "Fulfilled" && bodyRequest.Status != "In Progress" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid status"})
	}

	user := c.Locals("user").(*structs.User)

	if bodyRequest.Status == "Voided" && user.Role != "admin" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "You are not authorized to update this order"})
	}

	if bodyRequest.Status == "Voided" && bodyRequest.VoidReason == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid void reason"})
	}

	mongodb := c.Locals("mongo").(*mongo.Database)
	collection := mongodb.Collection("orders")

	order := new(structs.Order)
	err = collection.FindOne(context.Background(), bson.M{"orderNumber": orderId}).Decode(&order)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Order not found"})
	}

	if order.Status == bodyRequest.Status {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Status is identical to the current status"})
	}

	if order.PlacedBy != user.Email && user.Role != "admin" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"message": "You are not authorized to update this order"})
	}

	result, err := collection.UpdateOne(context.Background(), bson.M{"orderNumber": orderId}, bson.M{"$set": bson.M{"status": bodyRequest.Status}})
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Status Already Set as " + bodyRequest.Status})
	}

	if result.ModifiedCount == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Order not found"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Order updated successfully"})
}

func validateSetStatusBody(body structs.SetStatusRequest) []*structs.ValidationError {
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

package orders

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"inventory-api/src/utils/structs"
	"strconv"
	"time"
)

func Place(c *fiber.Ctx) error {
	bodyRequest := new(structs.Order)
	err := c.BodyParser(bodyRequest)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid request"})
	}

	errors := validatePlaceOrderBody(*bodyRequest)
	if len(errors) > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "Invalid JSON fields", "errors": errors})
	}
	var newName string
	for _, char := range bodyRequest.CustomerName {
		if char != ' ' {
			newName += string(char)
		} else {
			newName += "_"
		}
	}
	bodyRequest.DatePlaced = time.Now().Format("01-02-2006")
	bodyRequest.OrderNumber = strconv.Itoa(int(time.Now().Month())) + "-" + strconv.Itoa(time.Now().Day()) + "-" + strconv.Itoa(time.Now().Year()) + "-" + newName
	bodyRequest.Status = "In Progress"

	mongodb := c.Locals("mongo").(*mongo.Database)
	collection := mongodb.Collection("orders")
	user := c.Locals("user").(*structs.User)

	_, err = collection.InsertOne(context.Background(), bson.M{
		"cloneId":      uuid.New().String(),
		"use":          bodyRequest.Use,
		"customerName": bodyRequest.CustomerName,
		"datePlaced":   bodyRequest.DatePlaced,
		"dateRequired": bodyRequest.DateRequired,
		"orderNumber":  bodyRequest.OrderNumber,
		"clones":       bodyRequest.Clones,
		"status":       bodyRequest.Status,
		"notes":        bodyRequest.Notes,
		"placedBy":     user.Email,
	})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error placing order"})
	}

	clones := bodyRequest.Clones
	for _, clone := range clones {
		collection := mongodb.Collection("clones")
		curClone := new(structs.Clone)
		err = collection.FindOne(context.Background(), bson.M{"cloneId": clone.CloneId}).Decode(curClone)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error finding selected clone"})
		}
		_, err = collection.UpdateOne(context.Background(), bson.M{"cloneId": clone.CloneId}, bson.M{"$set": bson.M{"quantity": curClone.Quantity - clone.Quantity}})
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error updating clone status"})
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Order placed successfully", "order": bodyRequest})
}

func validatePlaceOrderBody(body structs.Order) []*structs.ValidationError {
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

package structs

import "github.com/google/uuid"

type Customer struct {
	MongoId        string `json:"id" bson:"_id"`
	CustomerId     string `json:"customerId" bson:"customerId"`
	Name           string `json:"name" bson:"name"`
	Email          string `json:"email" bson:"email"`
	SubscriptionId string `json:"subscriptionId" bson:"subscriptionId"`
	Phone          string `json:"phone" bson:"phone"`
	Address        string `json:"address" bson:"address"`
	City           string `json:"city" bson:"city"`
	State          string `json:"state" bson:"state"`
	Zip            string `json:"zip" bson:"zip"`
	Notes          string `json:"notes" bson:"notes"`
}

func (customer *Customer) GenerateCustomerId() {
	customer.CustomerId = uuid.New().String()
}

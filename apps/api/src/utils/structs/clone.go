package structs

type Clone struct {
	MongoId  string  `json:"id" bson:"_id"`
	CloneId  string  `json:"cloneId" bson:"cloneId"`
	Name     string  `json:"name" bson:"name"`
	Quantity float64 `json:"quantity" bson:"quantity"`
	Date     string  `json:"date" bson:"date"`
}

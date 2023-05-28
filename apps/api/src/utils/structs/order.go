package structs

type Order struct {
	MongoId      string  `json:"id" bson:"_id"`
	Use          string  `json:"use" bson:"use"`
	CustomerName string  `json:"customerName" bson:"customerName"`
	DatePlaced   string  `json:"datePlaced" bson:"datePlaced"`
	DateRequired string  `json:"dateRequired" bson:"dateRequired"`
	OrderNumber  string  `json:"orderNumber" bson:"orderNumber"`
	Clones       []Clone `json:"clones" bson:"clones"`
	Status       string  `json:"status" bson:"status"`
	Notes        string  `json:"notes" bson:"notes"`
	PlacedBy     string  `json:"placedBy" bson:"placedBy"`
}

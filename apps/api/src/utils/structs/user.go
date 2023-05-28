package structs

import "github.com/google/uuid"

type User struct {
	MongoId   string `json:"id" bson:"_id"`
	UserId    string `json:"userId" bson:"userId"`
	APIKey    string `json:"apikey" bson:"apikey"`
	Email     string `json:"email" bson:"email"`
	FirstName string `json:"firstName" bson:"firstName"`
	LastName  string `json:"lastName" bson:"lastName"`
	Password  string `json:"password" bson:"password"`
	Role      string `json:"role" bson:"role"`
}

func (user *User) Clean() {
	user.Password = ""
}

func (user *User) GenerateAPIKey() {
	user.APIKey = uuid.New().String()
}

func (user *User) GenerateUserID() {
	user.UserId = uuid.New().String()
}

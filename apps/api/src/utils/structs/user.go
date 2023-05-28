package structs

import "github.com/google/uuid"

type User struct {
	MongoId  string `json:"id" bson:"_id"`
	UserId   string `json:"userId" bson:"userId"`
	APIKey   string `json:"apikey" bson:"apikey"`
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
	Admin    bool   `json:"admin" bson:"admin"`
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

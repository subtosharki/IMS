package database

import (
	"context"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

var Client *mongo.Client

func Connect() {
	possibleEnv := os.Getenv("MONGO_URL")
	var envFile map[string]string
	var clientOptions *options.ClientOptions
	if possibleEnv == "" {
		var err error
		envFile, err = godotenv.Read("../.env")
		if err != nil {
			panic(err)
		}
		clientOptions = options.Client().ApplyURI(envFile["MONGO_URL"])
	} else {
		clientOptions = options.Client().ApplyURI(possibleEnv)
	}

	var err error
	Client, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		panic(err)
	}

	println("Connected to MongoDB!")
}

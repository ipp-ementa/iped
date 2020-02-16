package db

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Db holds the state of the database connection
var Db *mongo.Database

// Initializes database once go program starts or package is imported
// Panics if database couldn't be open and created
func init() {

	conn := os.Getenv("IPEW_CONNECTION_STRING")

	client, err := mongo.NewClient(options.Client().ApplyURI(conn))

	if err != nil {

		panic("Could not create MongoDB client")

	}

	ctx, cancelFunction := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	defer cancelFunction()

	if err != nil {

		panic("Could not connect to MongoDB server")

	}

	database := client.Database("ipew")

	Db = database

}

package database

import (
	"challeng-bravo/src/config"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Db() *mongo.Client {
	config.Loader()

	clientOptions := options.Client().ApplyURI(config.StringConnectionDB)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)

	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
	return client
}

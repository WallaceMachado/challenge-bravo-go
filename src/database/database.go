package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Db() *mongo.Client {

	clientOptions := options.Client().ApplyURI("mongodb+srv://admin:root@cluster0.pamgw.mongodb.net") // Connect to //MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("erro", err)
	}
	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		fmt.Println("erro", err)
	}
	fmt.Println("Connected to MongoDB!")
	return client
}

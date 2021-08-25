package database

import (
	"challeng-bravo/src/config"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetCollection(collection string) (*mongo.Collection, error) {

	client, err := mongo.NewClient(options.Client().ApplyURI(config.StringConnectionDB))

	if err != nil {
		fmt.Println("erro1")
		return nil, err
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	if err != nil {
		fmt.Println("erro2")
		return nil, err
	}

	return client.Database(config.DB_Name).Collection(collection), nil
}

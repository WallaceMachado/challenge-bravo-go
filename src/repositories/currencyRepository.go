package repositories

import (
	"challeng-bravo/src/database"
	"challeng-bravo/src/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func ListAll() ([]models.Currency, error) {
	var collection, err = database.GetCollection("currencies")
	var ctx = context.Background()

	var currencies []models.Currency

	filter := bson.D{}

	cur, err := collection.Find(ctx, filter)

	if err != nil {
		return nil, err
	}

	for cur.Next(ctx) {

		var currency models.Currency
		err = cur.Decode(&currency)

		if err != nil {
			return nil, err
		}

		currencies = append(currencies, currency)
	}

	return currencies, nil
}

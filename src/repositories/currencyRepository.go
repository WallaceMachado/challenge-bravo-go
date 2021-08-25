package repositories

import (
	"challeng-bravo/src/database"
	"challeng-bravo/src/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

var currencyCollection = database.Db().Database("chBravoDb").Collection("currencies") // get collection "currencies" from db() which returns *mongo.Client

func ListAll() ([]models.Currency, error) {
	var currencies []models.Currency
	result, err := currencyCollection.Find(context.TODO(), bson.D{{}}) //returns a *mongo.Cursor
	if err != nil {

		return nil, err

	}
	for result.Next(context.TODO()) { //Next() gets the next document for corresponding cursor

		var currency models.Currency
		err := result.Decode(&currency)
		if err != nil {
			return nil, err
		}

		currencies = append(currencies, currency) // appending document pointed by Next()
	}
	result.Close(context.TODO()) // close the cursor once stream of documents has exhausted

	return currencies, nil
}

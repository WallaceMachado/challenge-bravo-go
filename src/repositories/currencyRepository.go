package repositories

import (
	"challeng-bravo/src/database"
	"challeng-bravo/src/models"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func Create(currency models.Currency) (interface{}, error) {

	timeNow := time.Now()
	currency.Created_at = timeNow
	currency.Updated_at = timeNow

	insertResult, err := currencyCollection.InsertOne(context.TODO(), currency)
	if err != nil {

		return nil, err

	}

	fmt.Println(insertResult.InsertedID)
	return insertResult, nil // return the //mongodb ID of generated document
}

// Atualizar altera as informações de um usuário no banco de dados
func Update(currency models.Currency, ID string) error {

	fmt.Println(currency.Name)

	_id, err := primitive.ObjectIDFromHex(ID)

	if err != nil {
		return err
	}

	filter := bson.M{"_id": _id} // converting value to BSON type
	// for returning updated document

	update := bson.M{"$set": bson.M{"name": currency.Name, "code": currency.Code, "valieInUSD": currency.ValueInUSD, "updated_at": time.Now()}}

	if _, err := currencyCollection.UpdateOne(context.TODO(), filter, update); err != nil {
		return err
	}

	return nil
}

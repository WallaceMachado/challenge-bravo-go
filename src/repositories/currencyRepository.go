package repositories

import (
	"challeng-bravo/src/database"
	"challeng-bravo/src/models"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var currencyCollection = database.Db().Database("chBravoDb").Collection("currencies")

func ListAll() ([]models.Currency, error) {
	var currencies []models.Currency
	result, err := currencyCollection.Find(context.TODO(), bson.D{{}})
	if err != nil {

		return nil, err

	}
	for result.Next(context.TODO()) {

		var currency models.Currency
		err := result.Decode(&currency)
		if err != nil {
			return nil, err
		}

		currencies = append(currencies, currency)
	}
	result.Close(context.TODO())

	return currencies, nil
}

func GetByCode(code string) (models.Currency, error) {
	var currency models.Currency

	filter := bson.M{"code": code}

	if err := currencyCollection.FindOne(context.TODO(), filter).Decode(&currency); err != nil {

		return currency, err

	}

	return currency, nil

}

func GetById(id string) (models.Currency, error) {
	var currency models.Currency

	filter := bson.M{"_id": id}

	if err := currencyCollection.FindOne(context.TODO(), filter).Decode(&currency); err != nil {

		return currency, err

	}

	return currency, nil

}

func Create(currency models.Currency) (interface{}, error) {

	timeNow := time.Now()
	currency.Created_at = timeNow
	currency.Updated_at = timeNow

	insertResult, err := currencyCollection.InsertOne(context.TODO(), currency)
	if err != nil {

		return nil, err

	}

	return insertResult, nil
}

func Update(currency models.Currency, ID string) error {

	_id, err := primitive.ObjectIDFromHex(ID)

	if err != nil {
		return err
	}

	filter := bson.M{"_id": _id}

	update := bson.M{"$set": bson.M{"name": currency.Name, "code": currency.Code, "valieInUSD": currency.ValueInUSD, "updated_at": time.Now()}}

	if _, err := currencyCollection.UpdateOne(context.TODO(), filter, update); err != nil {
		return err
	}

	return nil
}

func Delete(ID string) error {

	_id, err := primitive.ObjectIDFromHex(ID)

	if err != nil {
		return err
	}

	filter := bson.M{"_id": _id}

	_, err = currencyCollection.DeleteOne(context.TODO(), filter)

	if err != nil {
		return err
	}

	return nil
}

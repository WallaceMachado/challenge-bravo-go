package cache

import (
	"challeng-bravo/src/config"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

var Client = redis.NewClient(&redis.Options{
	Addr:     config.AddrRedis,
	Password: config.PassRedis,
	DB:       0,
})

var ctx = context.TODO()

func Save(key string, value interface{}, expiryTimeInSeconds time.Duration) {

	var inInterface map[string]interface{}
	inrec, _ := json.Marshal(value)
	json.Unmarshal(inrec, &inInterface)

	err := Client.HSet(ctx, key, inInterface).Err()

	if err != nil {

		fmt.Println(err)
	}

	Client.Expire(ctx, key, expiryTimeInSeconds*time.Second)

}

func Recover(key string) (map[string]string, error) {

	data, err := Client.HGetAll(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	return data, nil

}

func Get(key string) (string, error) {

	data, err := Client.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}

	return data, nil

}

func SaveSet(key string, value interface{}, expiryTimeInSeconds time.Duration) {

	valueInByte, _ := json.Marshal(value)

	err := Client.Set(ctx, key, valueInByte, expiryTimeInSeconds*time.Second).Err()

	if err != nil {

		fmt.Println(err)
	}

}

func Delete(key string) {

	err := Client.Del(ctx, key)
	if err != nil {
		fmt.Println(err)
	}

}

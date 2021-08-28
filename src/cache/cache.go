package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

var Client = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "",
	DB:       0,
})

var ctx = context.TODO()

func Save(key string, value interface{}, expiryTimeInSeconds time.Duration) {

	var inInterface map[string]interface{}
	inrec, _ := json.Marshal(value)
	json.Unmarshal(inrec, &inInterface)

	err := Client.HSet(ctx, key, inInterface).Err()

	Client.Expire(ctx, key, expiryTimeInSeconds*time.Second)

	if err != nil {

		fmt.Println(err)
	}

}

func Recover(key string) (map[string]string, error) {

	data, err := Client.HGetAll(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	return data, nil

}

func Delete(key string) {

	err := Client.Del(ctx, key)
	if err != nil {
		fmt.Println(err)
	}

}

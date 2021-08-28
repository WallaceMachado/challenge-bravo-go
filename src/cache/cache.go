package cache

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

var Client = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "",
	DB:       0,
})

func Save(key string, value interface{}, expiryTimeInSeconds time.Duration) {
	expiryTimeInNanoSeconds := expiryTimeInSeconds * 1000000000

	valueMarshel, err := json.Marshal(value)
	if err != nil {

		fmt.Println(err)
	}
	err = Client.Set(key, valueMarshel, expiryTimeInNanoSeconds).Err()
	if err != nil {

		fmt.Println(err)
	}

}

func Recover(key string) (interface{}, error) {

	data, err := Client.Get(key).Result()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func Delete(key string) {

	err := Client.Del(key)
	if err != nil {
		fmt.Println(err)
	}

}

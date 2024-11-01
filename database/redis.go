package database

import (
	"context"
	"data_app/api"
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
)

var RCtx = context.Background()
var RDB *redis.Client

func InitRedis() {
	RDB = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", api.Config.Redis.Host, api.Config.Redis.Port),
		Password: api.Config.Redis.Password,
		DB:       api.Config.Redis.DB,
	})

	err := TestRedisConnection()
	if err != nil {
		log.Fatalf("Error while connecting redis: %v", err)
	} else {
		fmt.Println("Redis connection successful")
	}

}

func TestRedisConnection() error {
	_, err := RDB.Ping(RCtx).Result()
	if err != nil {
		return err
	}
	return nil
}

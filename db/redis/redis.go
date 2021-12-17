package redis_service

import (
	"fmt"

	"github.com/go-redis/redis"
)

func InitAndGetDb(dbUrl string) (*redis.Client) {
	fmt.Println("Go Redis Tutorial")

	client := redis.NewClient(&redis.Options{
		Addr: dbUrl,
		Password: "",
		DB: 0,
	})


	return client
}
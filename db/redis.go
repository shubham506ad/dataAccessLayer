package db

import (
	"fmt"

	"github.com/go-redis/redis"
)

type redisClient struct {
	cl *redis.Client
}

func NewRedisClient(dbUrl string) DbConnector {
	client := redis.NewClient(&redis.Options{
		Addr: dbUrl,
		Password: "",
		DB: 0,
	})
	return &redisClient{cl: client}
}

func (rc *redisClient) Connect() error {
	result,_ := rc.cl.Ping().Result()
	fmt.Println("connected to redis", result)
	return nil
}
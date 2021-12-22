package db

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
)

type redisClient struct {
	cl *redis.Client
}

func newRedisClient(dbUrl string) DbConnector {
	cl := redis.NewClient(&redis.Options{
		Addr: dbUrl,
	})
	return &redisClient{cl: cl}
}

func (rc *redisClient) Connect() error {
	_, err := rc.cl.Ping(context.TODO()).Result()
	if err != nil {
		return err
	}
	return nil
}

func (rc *redisClient) FindOne(ctx context.Context, collection string, filter interface{}) (interface{}, error) {
	val, err := rc.cl.Get(context.TODO(), collection).Result()
	if err != nil {
		return nil, err
	}
	return val, nil
}

func (rc *redisClient) FindMany(ctx context.Context, collection string, filter interface{}) (interface{}, error) {
	collectionSplit := strings.Split(collection, ",")
	return rc.cl.MGet(context.TODO(), collectionSplit...).Result()

}

func (rc *redisClient) InsertOne(ctx context.Context, collection string, document interface{}) (interface{}, error) {
	exp := time.Duration(600 * time.Second) // 10 minutes
	data := rc.cl.Set(context.TODO(), collection, document, exp).Err()
	return data, nil
}

func (rc *redisClient) InsertMany(ctx context.Context, collection string, document []interface{}) (interface{}, error) {

	exp := time.Duration(600 * time.Second) // 10 minutes
	var ifaces []interface{}
	pipe := rc.cl.TxPipeline()
	collectionSplit := strings.Split(collection, ",")

	for i := range collectionSplit {
		ifaces = append(ifaces, collectionSplit[i], document[i])
		pipe.Expire(context.TODO(), collectionSplit[i], exp)
	}

	if err := rc.cl.MSet(context.TODO(), ifaces...).Err(); err != nil {
		return nil, err
	}
	if _, err := pipe.Exec(context.TODO()); err != nil {
		return nil, err
	}
	return nil, nil
}

func (rc *redisClient) UpdateOne(ctx context.Context, collection string, filter interface{}, update interface{}) (interface{}, error) {
	return nil, nil
}

func (rc *redisClient) UpdateMany(ctx context.Context, collection string, filter interface{}, update interface{}) (interface{}, error) {
	return nil, nil
}

func (rc *redisClient) Cancel() error {
	client := rc.cl
	if client == nil {
		return nil
	}
	err := client.Close()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connection to redis closed.")
	return nil
}

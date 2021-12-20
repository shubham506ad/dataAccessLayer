package db

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type StorageType int

const (
    mongoDB StorageType = 1 << iota
    redisDB
)

type DbConnector interface {
	Connect() error
}

func NewStore(t StorageType, DBurl string) DbConnector {
    switch t {
    case mongoDB:
		return NewMongoClient(DBurl)
    case redisDB:
		return NewRedisClient(DBurl)
	}
	return nil
}

type StoreHelper interface {
	Connect() error
}

type mongoClient struct {
	cl *mongo.Client
}

func NewMongoClient(dbUrl string) StoreHelper {
	c, err := mongo.NewClient(options.Client().ApplyURI(dbUrl))
	if err != nil {
		panic(err)
	}
	return &mongoClient{cl: c}

}

func (mc *mongoClient) Connect() error {
	ctx, _ := context.WithTimeout(context.Background(), 30 * time.Second)
	err := mc.cl.Connect(ctx)
	if err != nil {
		return err
	}
	err = mongoPing(mc.cl, ctx)
	if err != nil {
		return err
	}
	return nil
}

type redisClient struct {
	cl *redis.Client
}

func NewRedisClient(dbUrl string) StoreHelper {
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

func mongoPing(client *mongo.Client, ctx context.Context) error{

	if err := client.Ping(ctx, readpref.Primary()); 
	err != nil {
		return err
	}
	fmt.Println("connected successfully")
	return nil
}
package db

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type mongoClient struct {
	cl *mongo.Client
}

func NewMongoClient(dbUrl string) DbConnector {
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

func mongoPing(client *mongo.Client, ctx context.Context) error{

	if err := client.Ping(ctx, readpref.Primary()); 
	err != nil {
		return err
	}
	fmt.Println("connected successfully")
	return nil
}
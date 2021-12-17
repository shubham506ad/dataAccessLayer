package mongo_service

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


type MongoDbHelper interface {
	Client() ClientHelper
}

type ClientHelper interface {
	Connect() error
}

type mongoClient struct {
	cl *mongo.Client
}
type mongoDatabase struct {
	db *mongo.Database
}

func NewClient(dbUrl string) ClientHelper {
	c, err := mongo.NewClient(options.Client().ApplyURI(dbUrl))
	if err != nil {
		panic(err)
	}
	return &mongoClient{cl: c}

}

func (mc *mongoClient) Connect() error {
	ctx, _ := context.WithTimeout(context.Background(), 30 * time.Second)
	return mc.cl.Connect(ctx)
}

func (md *mongoDatabase) Client() ClientHelper {
	client := md.db.Client()
	return &mongoClient{cl: client}
}

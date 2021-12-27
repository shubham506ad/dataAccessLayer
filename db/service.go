// provides an interface for the accessing multiple databases
package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

type StorageType int

const (
    mongoDB StorageType = 1 << iota
    redisDB
)

/*
The functions that are exposed to be used by multiple databases
*/
type DbConnector interface {
	Connect() error
	FindOne(context.Context, string, interface{}) (interface{}, error)
	FindMany(context.Context, string, interface{}) ([]bson.M, error)
	InsertOne(context.Context, string, interface{}) (interface{}, error)
	InsertMany(context.Context, string, []interface{}) ([]interface{}, error)
	UpdateOne(context.Context, string, interface{}, interface{}) (interface{}, error)
	UpdateMany(context.Context,string, interface{}, interface{}) (interface{}, error)
	Cancel() error
}

type SingleResultHelper interface {
	Decode(v interface{}) error
}

func NewStore(t StorageType, DBurl string, DBname string) DbConnector {
    switch t {
    case mongoDB:
		return newMongoClient(DBurl, DBname)
    case redisDB:
		return nil
	}
	return nil
}
package db

import "context"

type StorageType int

const (
    mongoDB StorageType = 1 << iota
    redisDB
)

type DbConnector interface {
	Connect() error
	FindOne(context.Context, string, interface{}) (interface{}, error)
	FindMany(context.Context, string, interface{}) (interface{}, error)
	InsertOne(context.Context, string, interface{}) (interface{}, error)
	InsertMany(context.Context, string, []interface{}) (interface{}, error)
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
package db

import (
	mongo_service "datalayer/db/mongo"
	"fmt"
)

type StorageType int

const (
    mongoDB StorageType = 1 << iota
    redisDB
)

type DbConnector interface {
	mongo_service.ClientHelper
}

func NewStore(t StorageType, DBurl string) DbConnector {
    switch t {
    case mongoDB:
		return mongo_service.NewClient(DBurl)
    case redisDB:
		fmt.Println("redis service will be called")
	}
	return nil
}
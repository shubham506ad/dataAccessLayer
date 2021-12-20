package db

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
		return newMongoClient(DBurl)
    case redisDB:
		return newRedisClient(DBurl)
	}
	return nil
}
package main

import (
	"datalayer/db"
	"fmt"
)

func main()  {
	mongoClient := db.NewStore(1,"mongodb://localhost:27017") 
	pong := mongoClient.Connect()
	fmt.Println("getting the value mongo", pong)
	redisClient := db.NewStore(2, "localhost:6379")
	redisPong := redisClient.Connect()
	fmt.Println("getting the value redis", redisPong)

}
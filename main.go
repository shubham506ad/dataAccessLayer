package main

import (
	"datalayer/db"
	"fmt"
)

func main()  {
	client := db.NewStore(1,"mongodb://localhost:27017") 
	pong := client.Connect()
	fmt.Println("getting the value", pong)
}
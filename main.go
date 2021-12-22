package main

import (
	"context"
	"datalayer/db"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

func main()  {
	mongoClient := db.NewStore(1,"mongodb://localhost:27017", "tester") 
	mongoClient.Connect()
	insertDoc := bson.D{{"title", "Record of a Shriveled Datum"}, {"text", "No bytes, no problem. Just insert a document, in MongoDB"}, {"index", 1}}
	insertOneRes, err := mongoClient.InsertOne(context.TODO(), "testCol", insertDoc)
	if err != nil {
		panic(err)
	}
	fmt.Println(insertOneRes)

	insertDocArr := []interface{}{
	bson.D{{"title", "Record of a Shriveled Datum"}, {"text", "No bytes, no problem. Just insert a document, in MongoDB"}, {"index", 2}},
	bson.D{{"title", "Showcasing a Blossoming Binary"}, {"text", "Binary data, safely stored with GridFS. Bucket the data"}, {"index", 3}},
	}
	insertMultiRes, err := mongoClient.InsertMany(context.TODO(), "testCol", insertDocArr)
	if err != nil {
		panic(err)
	}
	fmt.Println(insertMultiRes)

	findMulDoc, err := mongoClient.FindMany(context.TODO(), "testCol", bson.M{})
	if err != nil {
		panic(err)
	}
	fmt.Println(findMulDoc)

	findSingleDoc, err := mongoClient.FindOne(context.TODO(), "testCol", bson.M{})
	if err != nil {
		panic(err)
	}
	fmt.Println(findSingleDoc)

	updateSingleDoc, err := mongoClient.UpdateOne( context.TODO(), "testCol", bson.M{"index": 1}, bson.D{ {"$set", bson.D{{"text", "things have changed now !!"}}}})
	if err != nil {
		panic(err)
	}
	fmt.Println(updateSingleDoc)
	
	updateMultiDoc, err := mongoClient.UpdateMany( context.TODO(), "testCol", bson.D{}, bson.D{ {"$set", bson.D{{"text", "things have changed now !!"}}}})
	if err != nil {
		panic(err)
	}
	fmt.Println(updateMultiDoc)
	
}
package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Person struct {
	Name string
	Age  int
}

func main() {

	// Set client options
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, e := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://foo:bar@localhost:27017"))

	CheckError(e)

	// // Check the connection
	e = client.Ping(context.TODO(), nil)
	// CheckError(e)

	// // get collection as ref
	// collection := client.Database("testdb").Collection("people")

	// // insert
	// john := Person{"John", 24}
	// jane := Person{"Jane", 27}
	// ben := Person{"Ben", 16}

	// _, e = collection.InsertOne(context.TODO(), john)
	// CheckError(e)

	// persons := []interface{}{jane, ben}
	// _, e = collection.InsertMany(context.TODO(), persons)
	// CheckError(e)

	// // update
	// filter := bson.D{{"name", "John"}}

	// update := bson.D{
	// 	{"$set", bson.D{
	// 		{"age", 26},
	// 	}},
	// }

	// _, e = collection.UpdateOne(context.TODO(), filter, update)
	// CheckError(e)

	// // find
	// var res Person
	// e = collection.FindOne(context.TODO(), filter).Decode(&res)
	// fmt.Println(res)

	// // delete
	// _, e = collection.DeleteMany(context.TODO(), bson.D{{}})
	// CheckError(e)
}

func CheckError(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

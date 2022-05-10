package db

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DbConstructor struct {
	client   *mongo.Client
	database string
}

var DatabaseInstance DbConstructor

const (
	username = ""
	password = ""
)

func (db *DbConstructor) Database() *mongo.Database {
	return db.client.Database(db.database)
}

func (db *DbConstructor) GetCollection(collection string) *mongo.Collection {
	return db.client.Database(db.database).Collection(collection)
}

func (db *DbConstructor) Setup() (client *mongo.Client) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URL")))

	if err != nil {
		panic(err)
	}

	db.database = os.Getenv("DATABASE_NAME")
	db.client = client

	return
}

func (db *DbConstructor) Shutdown() error {
	return db.client.Disconnect(context.TODO())
}

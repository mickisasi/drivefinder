package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	client *mongo.Client
}

func New(dbUrl string) (*Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dbUrl))
	if err != nil {
		return nil, err
	}

	fmt.Println("connected to database")
	return &Database{client: client}, nil
}

func (db *Database) OpenCollection(collectionName string) *mongo.Collection {
	return db.client.Database("drivefinder").Collection(collectionName)
}

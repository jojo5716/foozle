package api

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func connect() {
	host := "localhost"
	port := 27017

	clientOpts := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%d", host, port))
	client, _ = mongo.Connect(context.TODO(), clientOpts)
}

func Initialize() {
	connect()

	fmt.Println("Congratulations, you're already connected to MongoDB!")
}

func GetClient() *mongo.Client {
	return client
}

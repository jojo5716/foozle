package api

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var dbClient *mongo.Client

const (
	DBName = "foozle"
)

type Connection interface {
	Connect() *mongo.Client
	GetCollection(string) *mongo.Collection
}

type Client struct {
	Host string
	Port int32
}

func (m *Client) Connect() *mongo.Client {
	clientOpts := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%d", m.Host, m.Port))
	dbClient, _ = mongo.Connect(context.TODO(), clientOpts)

	return dbClient
}

func (m *Client) GetCollection(collectionName string) *mongo.Collection {
	collection := dbClient.Database(DBName).Collection(collectionName)

	return collection
}

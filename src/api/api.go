package api

import (
	"fmt"
	// "go.mongodb.org/mongo-driver/mongo"
)

// var dbClient *mongo.Client


type Connection interface {
	Connect() string
}

type Client struct {
	Host string
	Port int32
}

func (m *Client) Connect() string {
	return m.Host
}

func Initialize() {
	client := Client{Host: "localhost", Port: 27017}

	var connection Connection
	connection = &client
	fmt.Println(connection.Connect())
}

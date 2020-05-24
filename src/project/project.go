package project

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"api"
)

var projectCollection *mongo.Collection

type Project struct {
	Uuid, Name string
	Connection api.Connection
}

func (p *Project) getCollection() *mongo.Collection {
	if projectCollection == nil {
		projectCollection = p.Connection.GetCollection("project")
	}
	return projectCollection
}

func (p *Project) Exist() bool {
	var result Project

	projectCollection := p.getCollection()

	filter := bson.D{{"uuid", p.Uuid}}

	projectCollection.FindOne(context.TODO(), filter).Decode(&result)

	return result.Uuid != ""
}

func (p *Project) Create() {
	projectCollection := p.getCollection()

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	projectCollection.InsertOne(ctx, p)
}

package project

import (
	"context"
	"log"
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
	var hasToCreateProject bool

	projectCollection := p.getCollection()

	err := projectCollection.FindOne(context.TODO(), bson.D{}).Decode(&result)

	if err != nil {
		log.Fatal("ERROR: %v\n", err)
		hasToCreateProject = false
	} else {
		hasToCreateProject = true
	}

	return hasToCreateProject

}

func (p *Project) Create() {
	projectCollection := p.getCollection()

	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	projectCollection.InsertOne(ctx, p)
}

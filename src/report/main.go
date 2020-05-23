package report

import (
	"fmt"

	// Internal packages
	"api"
	"project"
)

const (
	DBName = "foozle"
	HOST   = "localhost"
	PORT   = 27017
)

func Create() {
	fmt.Println("Creating report ...")
	client := api.Client{Host: HOST, Port: PORT}
	var connection api.Connection
	connection = &client

	connection.Connect()

	p := project.Project{
		Uuid:       "22",
		Name:       "ee",
		Connection: connection,
	}

	if !p.Exist() {
		p.Create()
	}
}

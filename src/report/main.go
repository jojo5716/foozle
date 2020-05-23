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

type VisitReport struct {
	Project     string      `json:"project"`
	LoadedOn    string      `json:loadedOn`
	SessionTemp string      `json:sessionTemp`
	Session     string      `json:session`
	PageToken   string      `json:pageToken`
	EventTime   float32     `json:eventTime`
	Data        interface{} `json:"data"`
	Page        interface{} `json:page`
	Enviroment  interface{} `json:enviroment`
	MetaData    interface{} `json:metaData`
	Actions     interface{} `json:actions`
	UserInfo    interface{} `json:userInfo`
}

func Create(visitReport VisitReport) {
	fmt.Println("Creating report ...", visitReport)

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

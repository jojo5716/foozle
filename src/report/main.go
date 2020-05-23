package report

import (
	"fmt"
	// Internal packages
	"api"
	"project"
)

func Create() {
	fmt.Println("Creating report ...")

	api.Initialize()

	p := project.Project{Uuid: "22", Name: "ee"}

	p.Create()
}

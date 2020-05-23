package project

import "fmt"

type Project struct {
	Uuid, Name string
}

func (p *Project) Create() {
	fmt.Println("Creating project...")
}

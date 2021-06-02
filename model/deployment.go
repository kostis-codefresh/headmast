package model

type Deployment struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Version  string `json:"version"`
	Endpoint string `json:"endpoint"`

	ParentApplication *Application
	ParentEnvironment *Environment
	Features          []*Feature `json:"-"`
}

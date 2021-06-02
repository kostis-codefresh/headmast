package model

type Application struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Version string `json:"version"`

	ActiveDeployments []*Deployment `json:"-"`
}

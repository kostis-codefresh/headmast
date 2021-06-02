package model

type Environment struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`

	ActiveDeployments []*Deployment
}

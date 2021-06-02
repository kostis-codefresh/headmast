package model

type Application struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`

	ActiveDeployments []*Deployment `json:"-"`
}

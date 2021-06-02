package model

type Feature struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Link        string `json:"link"`
	Description string `json:"description"`

	ParentDeployments []*Deployment
}

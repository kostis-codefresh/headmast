package server

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/kostis-codefresh/headmast/model"
)

type product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func (p *product) getProduct(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (p *product) updateProduct(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (p *product) deleteProduct(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (p *product) createProduct(db *sql.DB) error {
	return errors.New("Not implemented")
}

func getProducts(db *sql.DB, start, count int) ([]product, error) {
	return nil, errors.New("Not implemented")
}

func (b *Backend) createEnvironment(w http.ResponseWriter, r *http.Request) {
	var env model.Environment
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&env); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	b.lastIdUsed++
	env.ID = b.lastIdUsed
	b.Environments = append(b.Environments, env)

	respondWithJSON(w, http.StatusCreated, env)
}

func (b *Backend) getEnvironments(w http.ResponseWriter, r *http.Request) {

	// if err != nil {
	// 	respondWithError(w, http.StatusInternalServerError, err.Error())
	// 	return
	// }

	respondWithJSON(w, http.StatusOK, b.Environments)
}

package server

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/kostis-codefresh/headmast/model"
)

func getApplication(db *sql.DB) (*model.Application, error) {
	return nil, errors.New("Not implemented")
}

func updateApplication(db *sql.DB) (*model.Application, error) {
	return nil, errors.New("Not implemented")
}

func deleteApplication(db *sql.DB) (*model.Application, error) {
	return nil, errors.New("Not implemented")
}

func createApplication(a *model.Application) (*model.Application, error) {
	return nil, errors.New("Not implemented")
}

func getApplications(db *sql.DB, start, count int) ([]model.Application, error) {
	return nil, errors.New("Not implemented")
}

func (b *Backend) createApplication(w http.ResponseWriter, r *http.Request) {
	var app model.Application
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&app); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	b.Applications = append(b.Applications, app)

	respondWithJSON(w, http.StatusCreated, app)
}

func (b *Backend) getApplications(w http.ResponseWriter, r *http.Request) {

	// if err != nil {
	// 	respondWithError(w, http.StatusInternalServerError, err.Error())
	// 	return
	// }

	respondWithJSON(w, http.StatusOK, b.Applications)
}

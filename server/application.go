package server

import (
	"encoding/json"
	"net/http"

	"github.com/kostis-codefresh/headmast/model"
)

func (b *Backend) createApplication(w http.ResponseWriter, r *http.Request) {
	var app model.Application
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&app); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	b.lastIdUsed++
	app.ID = b.lastIdUsed
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

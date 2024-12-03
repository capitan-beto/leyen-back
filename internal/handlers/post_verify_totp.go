package handlers

import (
	"cmd/api/main.go/api"
	"cmd/api/main.go/internal/tools"
	"cmd/api/main.go/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func PostVerifyTOTP(w http.ResponseWriter, r *http.Request) {
	var reqBody *models.VerifyTOTP

	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	defer r.Body.Close()

	var database *sql.DB
	database, err = tools.CreateConnection()
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
	}

	err = tools.VerifyTOTP(reqBody.UserEmail, reqBody.Totp, database)
	if err != nil {
		if err == fmt.Errorf("error: totp doesn't match") {
			log.Error(err)
			api.UnauthorizedHandler(w, err)
			return
		}
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var res = api.PostHandlerResponse{
		Code:    http.StatusOK,
		Message: fmt.Sprintf("Welcome!, %s", reqBody.UserEmail),
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	if err = json.NewEncoder(w).Encode(res); err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
}

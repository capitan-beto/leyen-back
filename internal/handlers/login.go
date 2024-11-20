package handlers

import (
	"cmd/api/main.go/api"
	"cmd/api/main.go/internal/tools"
	"cmd/api/main.go/models"
	"cmd/api/main.go/utils"
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	var user *models.User
	json.NewDecoder(r.Body).Decode(&user)

	var database *sql.DB
	database, err := tools.CreateConnection()
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	role, err := tools.AuthenticateUser(user.Email, user.Pass, database)
	if err != nil {
		log.Error(err)
		api.UnauthorizedHandler(w, errors.New("invalid credentials"))
	}

	tokenString, err := utils.GenerateToken(user.Email, *role)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
	}

	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(tokenString); err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
	}
}

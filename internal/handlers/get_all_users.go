package handlers

import (
	"cmd/api/main.go/api"
	"cmd/api/main.go/internal/tools"
	"cmd/api/main.go/models"
	"database/sql"
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	var database *sql.DB
	database, err := tools.CreateConnection()
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var usersDetails []*models.User
	usersDetails, err = tools.GetUsers(database)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var res = api.GetAllUsersResponse{
		Code:  http.StatusOK,
		Users: usersDetails,
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
}

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

func PostSignUp(w http.ResponseWriter, r *http.Request) {
	var database *sql.DB

	database, err := tools.CreateConnection()
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var newUser models.User
	json.NewDecoder(r.Body).Decode(&newUser)
	defer r.Body.Close()

	if err := tools.AddUser(&newUser, database); err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var res = api.PostHandlerResponse{
		Code:    http.StatusOK,
		Message: fmt.Sprintf("welcome, %v", newUser.FullName),
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	if err = json.NewEncoder(w).Encode(res); err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

}

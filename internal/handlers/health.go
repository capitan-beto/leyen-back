package handlers

import (
	"cmd/api/main.go/api"
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func Health(w http.ResponseWriter, r *http.Request) {

	var res = api.HealthResponse{
		Code:    http.StatusOK,
		Message: "Welcome!",
	}
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

}

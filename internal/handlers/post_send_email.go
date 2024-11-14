package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"cmd/api/main.go/api"
	"cmd/api/main.go/internal/tools"
	"cmd/api/main.go/models"

	log "github.com/sirupsen/logrus"
)

func PostSendEmail(w http.ResponseWriter, r *http.Request) {
	var reqBody *models.EmailReqBody
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		log.Error(err)
		api.RequestErrorHandler(w, err)
		return
	}

	to := strings.Split(reqBody.ToAddr, ",")

	err = tools.SendEmail(to, reqBody.Subj, reqBody.Body)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var res = api.PostHandlerResponse{
		Code:    http.StatusOK,
		Message: fmt.Sprintf("email sent to: %s", to),
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	if err = json.NewEncoder(w).Encode(res); err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

}

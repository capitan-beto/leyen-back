package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"cmd/api/main.go/api"
	"cmd/api/main.go/internal/tools"
	"cmd/api/main.go/models"

	log "github.com/sirupsen/logrus"
	"github.com/xlzd/gotp"
)

func PostAccountRecovery(w http.ResponseWriter, r *http.Request) {
	var reqBody *models.OTTPRequest

	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		log.Error(err)
		api.RequestErrorHandler(w, err)
		return
	}

	defer r.Body.Close()

	totp := gotp.NewDefaultTOTP(gotp.RandomSecret(16))

	var database *sql.DB
	database, err = tools.CreateConnection()
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
	}

	err = tools.UpdateLastTOTP(reqBody.ToAddr, totp.Now(), database)
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
	}

	to := strings.Split(reqBody.ToAddr, ",")

	err = tools.SendEmail(to, "Login con código único", fmt.Sprintf("<a>%s</a>", totp.Now()))
	if err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}

	var res = api.PostHandlerResponse{
		Code:    http.StatusOK,
		Message: fmt.Sprintf("one tiime code sent to %s", to),
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	if err = json.NewEncoder(w).Encode(res); err != nil {
		log.Error(err)
		api.InternalErrorHandler(w)
		return
	}
}

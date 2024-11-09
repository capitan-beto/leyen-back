package api

import (
	"cmd/api/main.go/models"
	"encoding/json"
	"net/http"
)

//error handling

type Error struct {
	Code    int
	Message string
}

func writeError(w http.ResponseWriter, msg string, code int) {
	res := Error{
		Code:    code,
		Message: msg,
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(res)
}

var (
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, err.Error(), http.StatusBadRequest)
	}
	InternalErrorHandler = func(w http.ResponseWriter) {
		writeError(w, "An unexpected Error Ocurred", http.StatusInternalServerError)
	}
	UnauthorizedHandler = func(w http.ResponseWriter, err error) {
		writeError(w, err.Error(), http.StatusUnauthorized)
	}
	NotFoundHandler = func(w http.ResponseWriter, err error) {
		writeError(w, err.Error(), http.StatusNotFound)
	}
)

// handlers responses

type HealthResponse struct {
	Code    int
	Message string
}

type GetAllUsersResponse struct {
	Code  int
	Users []*models.User
}

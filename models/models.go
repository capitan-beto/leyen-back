package models

type User struct {
	Dni          int    `json:"dni"`
	Email        string `json:"email"`
	FullName     string `json:"fullName"`
	Points       int    `json:"points"`
	RegisterDate string `json:"register_date"`
}

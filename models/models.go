package models

type User struct {
	Id           int    `json:"id"`
	Pass         string `json:"password"`
	Role         string `json:"role"`
	Email        string `json:"email"`
	FullName     string `json:"fullName"`
	Points       int    `json:"points"`
	RegisterDate string `json:"register_date"`
}

type OTTPRequest struct {
	ToAddr string `json:"user_email"`
}

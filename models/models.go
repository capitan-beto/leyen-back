package models

type User struct {
	Dni          int    `json:"dni"`
	Pass         string `json:"password"`
	Role         string `json:"role"`
	Email        string `json:"email"`
	FullName     string `json:"fullName"`
	Points       int    `json:"points"`
	RegisterDate string `json:"register_date"`
}

type EmailReqBody struct {
	ToAddr string `json:"to_addr"`
	Subj   string `json:"subject"`
	Body   string `json:"body"`
}

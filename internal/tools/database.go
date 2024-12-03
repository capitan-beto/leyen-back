package tools

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"cmd/api/main.go/models"
	"cmd/api/main.go/utils"

	"github.com/go-sql-driver/mysql"
	"github.com/xlzd/gotp"
)

func CreateConnection() (*sql.DB, error) {
	cfg := mysql.Config{
		User:                 os.Getenv("DB_USER"),
		Passwd:               os.Getenv("DB_PASS"),
		Net:                  "tcp",
		Addr:                 os.Getenv("DB_ADDR"),
		DBName:               os.Getenv("DB_NAME"),
		AllowNativePasswords: true,
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, err
	}

	pingErr := db.Ping()
	if pingErr != nil {
		return nil, err
	}

	db.SetMaxOpenConns(5)
	fmt.Println("Connected!")
	return db, nil
}

//users queries

func GetUsers(db *sql.DB) ([]*models.User, error) {
	var users []*models.User

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.Id, &u.Pass, &u.Role, &u.Email, &u.FullName, &u.Points, &u.RegisterDate, &u.LastTOTP); err != nil {
			return nil, err
		}
		users = append(users, &u)

	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	db.Close()
	return users, nil
}

// REMEMBER TO CHECK DUPLICATES

func AddUser(nu *models.User, db *sql.DB) error {
	totp := gotp.NewDefaultTOTP(gotp.RandomSecret(16))

	pass, err := utils.HashPassword(nu.Pass)
	if err != nil {
		return err
	}

	var query string = "INSERT INTO users (dni, pword, email, role, full_name, points, register_date, last_totp) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
	_, err = db.Exec(query, &nu.Id, &pass, &nu.Email, &nu.Role, &nu.FullName, &nu.Points, &nu.RegisterDate, totp.Now())
	if err != nil {
		return err
	}

	db.Close()
	return nil
}

func AuthenticateUser(email, pword string, db *sql.DB) (*string, error) {
	userCreds := struct {
		pword string
		role  string
	}{}

	row := db.QueryRow("SELECT pword, role FROM users WHERE email = ?", email)
	err := row.Scan(&userCreds.pword, &userCreds.role)
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, err
		}
		return nil, err
	}

	if err := utils.CheckPassword(userCreds.pword, pword); err != nil {
		return nil, err
	}

	db.Close()
	return &userCreds.role, err
}

func UpdateLastTOTP(email, totp, secret string, db *sql.DB) error {
	var err error

	query := "UPDATE users SET secret_totp = ? WHERE email = ?"
	_, err = db.Exec(query, totp, secret, email)
	if err != nil {
		return err
	}

	db.Close()
	return err
}

func VerifyTOTP(email, token string, db *sql.DB) error {
	var secret string

	row := db.QueryRow("SELECT secret_totp FROM users WHERE email = ?", email)
	err := row.Scan(&secret)
	if err != nil {
		if err == sql.ErrNoRows {
			return err
		}
		return fmt.Errorf("error: no such email, %v", err)
	}

	totp := gotp.NewDefaultTOTP(secret)

	currentUnixTime := time.Now().Unix()

	// time window
	const step = 30
	const window = 1 // Allow 1 step before and after

	// Check tokens for current time and within the window
	for i := -window; i <= window; i++ {
		timestamp := currentUnixTime + int64(i*step)
		if totp.Verify(token, timestamp) {
			return nil
		} else {
			err = fmt.Errorf("invalid token")
		}
	}

	return err
}

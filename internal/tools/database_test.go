package tools

import (
	"fmt"
	"regexp"
	"testing"

	"cmd/api/main.go/models"
	"cmd/api/main.go/utils"

	"github.com/DATA-DOG/go-sqlmock"
)

var testdata = models.User{
	Id:           1,
	Pass:         "carlosnana1",
	Role:         "admin",
	FullName:     "Carlos Nana",
	Points:       10038,
	RegisterDate: "01-10-2024",
}

func TestGetUsers(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create sqlmock: %v", err)
	}

	defer db.Close()

	rows := sqlmock.NewRows([]string{"dni", "password", "role", "email", "fullName", "points", "register_date"}).
		AddRow(40616528, "carlosnana1", "admin", "crnana98@gmail.com", "Carlos Nana", 10038, "01-10-2024")
	mock.ExpectQuery("SELECT(.*)").WillReturnRows(rows)

	res, err := GetUsers(db)
	if err != nil {
		t.Error(err)
	}

	got := models.User{
		Id:           res[0].Id,
		Pass:         res[0].Pass,
		Role:         res[0].Role,
		FullName:     res[0].FullName,
		Points:       res[0].Points,
		RegisterDate: res[0].RegisterDate,
	}

	if got != testdata {
		t.Fatalf("Expected 'Carlos Nana', got %v", res[0].FullName)
	}
}

func TestAddUsersError(t *testing.T) {
	const sqlInsert = `INSERT INTO users (dni, pword, email, role, full_name, points, register_date) VALUES
							(?, ?, ?, ?, ?, ?, ?)`

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create sqlmock: %v", err)
	}

	mock.ExpectExec(regexp.QuoteMeta(sqlInsert)).WillReturnError(fmt.Errorf("error code: 1062"))

	if err := AddUser(&testdata, db); err == nil {
		t.Fatalf("expected error code: 1062, got %v", err)
	}
}

func TestAddUserOk(t *testing.T) {
	const sqlInsert = `INSERT INTO users (dni, pword, email, role, full_name, points, register_date) VALUES
	(?, ?, ?, ?, ?, ?, ?)`

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create sqlmock: %v", err)
	}

	mock.ExpectExec(regexp.QuoteMeta(sqlInsert)).WillReturnResult(sqlmock.NewResult(1, 1))

	if err := AddUser(&testdata, db); err != nil {
		t.Fatal(err)
	}
}

func TestAuthenticateUserOk(t *testing.T) {
	hashedPwd, err := utils.HashPassword("supersecret1234")
	if err != nil {
		t.Fatal(err)
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	rows := sqlmock.NewRows([]string{"pword", "role"}).
		AddRow(hashedPwd, "admin")
	mock.ExpectQuery("SELECT(.*)").WillReturnRows(rows)

	role, err := AuthenticateUser("email@gmail.com", "supersecret1234", db)
	if err != nil {
		t.Fatal(err)
	}

	expected := "admin"
	if got := role; string(*got) != expected {
		t.Fatalf("error: expected %v, got %v", expected, string(*got))
	}
}

func TestAuthenticateUserAuthError(t *testing.T) {
	hashedPwd, err := utils.HashPassword("supersecret1234")
	if err != nil {
		t.Fatal(err)
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	rows := sqlmock.NewRows([]string{"pword", "role"}).
		AddRow(hashedPwd, "admin")
	mock.ExpectQuery("SELECT(.*)").WillReturnRows(rows)

	_, err = AuthenticateUser("email@gmail.com", "falsesupersecret1234", db)
	if err == nil {
		t.Fatal(err)
	}
}

func TestAuthenticateUserRoleError(t *testing.T) {
	hashedPwd, err := utils.HashPassword("supersecret1234")
	if err != nil {
		t.Fatal(err)
	}

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}

	defer db.Close()

	rows := sqlmock.NewRows([]string{"pword", "role"}).
		AddRow(hashedPwd, "user")
	mock.ExpectQuery("SELECT(.*)").WillReturnRows(rows)

	role, err := AuthenticateUser("email@gmail.com", "supersecret1234", db)
	if err != nil {
		t.Fatal(err)
	}

	expected := "admin"
	if got := role; string(*got) == expected {
		t.Fatalf("error: expected %v, got %v", expected, string(*got))
	}
}

func TestUpdateLastOTP(t *testing.T) {
	const mockQuery = "UPDATE users SET last_totp = ? WHERE email = ?"

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create sqlmock: %v", err)
	}

	mock.ExpectExec(regexp.QuoteMeta(mockQuery)).WillReturnResult(sqlmock.NewResult(1, 1))

	if err := UpdateLastTOTP("test@gmail.com", "123456", db); err != nil {
		t.Fatal(err)
	}
}

package tools

import (
	"testing"

	"cmd/api/main.go/models"

	"github.com/DATA-DOG/go-sqlmock"
)

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

	expected := models.User{
		Dni:          40616528,
		Pass:         "carlosnana1",
		Role:         "admin",
		FullName:     "Carlos Nana",
		Points:       10038,
		RegisterDate: "01-10-2024",
	}

	got := models.User{
		Dni:          res[0].Dni,
		Pass:         res[0].Pass,
		Role:         res[0].Role,
		FullName:     res[0].FullName,
		Points:       res[0].Points,
		RegisterDate: res[0].RegisterDate,
	}

	if got != expected {
		t.Fatalf("Expected 'Carlos Nana', got %v", res[0].FullName)
	}
}

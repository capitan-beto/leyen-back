package tools

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestGetUsers(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create sqlmock: %v", err)
	}

	defer db.Close()

	rows := sqlmock.NewRows([]string{"dni", "email", "fullName", "points", "register_date"}).
		AddRow(40616528, "crnana98@gmail.com", "Carlos Nana", 10038, "01-10-2024")
	mock.ExpectQuery("SELECT(.*)").WillReturnRows(rows)

	res, err := GetUsers(db)
	if err != nil {
		t.Error(err)
	}

	if res[0].FullName != "Carlos Nana" {
		t.Fatalf("Expected 'Carlos Nana', got %v", res[0].FullName)
	}
}

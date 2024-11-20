package utils

import "testing"

func TestHashPasswordOK(t *testing.T) {
	pass := "test"

	hashedPwd, err := HashPassword(pass)
	if err != nil {
		t.Error(err)
	}

	if err := CheckPassword(hashedPwd, pass); err != nil {
		t.Error(err)
	}
}

func TestCheckPasswordFalse(t *testing.T) {
	hashedPwd, err := HashPassword("test")
	if err != nil {
		t.Fatal(err)
	}

	if err := CheckPassword(hashedPwd, "thisshouldmatch"); err == nil {
		t.Fatalf("error: expected false, got %t", err)
	}
}

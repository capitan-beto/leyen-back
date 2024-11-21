package utils

import "testing"

func TestGenerateTokenAndVerifyToken(t *testing.T) {
	token, err := GenerateToken("test@gmail.com", "admin")
	if err != nil {
		t.Fatal(err)
	}

	if err = VerifyToken(token, "admin"); err != nil {
		t.Fatal(err)
	}
}

func TestGenerateTokenAndVerifyTokenErr(t *testing.T) {
	token, err := GenerateToken("test@gmail.com", "")
	if err != nil {
		t.Fatal()
	}

	if err = VerifyToken(token, "client"); err != nil {
		t.Fatal(err)
	}
}

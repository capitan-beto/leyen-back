package utils

import "testing"

func TestHashPasswordOK(t *testing.T) {
	pass := "test"

	hashedPwd, err := HashPassword(pass)
	if err != nil {
		t.Error(err)
	}

	if pwdMatch := CheckPassword(hashedPwd, pass); !pwdMatch {
		t.Error(err)
	}
}

func TestCheckPasswordFalse(t *testing.T) {
	hashedPwd, err := HashPassword("test")
	if err != nil {
		t.Fatal(err)
	}

	if pwdMatch := CheckPassword(hashedPwd, "thisshouldmatch"); pwdMatch {
		t.Fatalf("error: expected false, got %t", pwdMatch)
	}
}

package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(pass string) (string, error) {
	var passBytes = []byte(pass)

	hashedPassBytes, err := bcrypt.
		GenerateFromPassword(passBytes, bcrypt.MinCost)

	return string(hashedPassBytes), err
}

func CheckPassword(hashedPassword, clientPass string) error {
	err := bcrypt.CompareHashAndPassword(
		[]byte(hashedPassword), []byte(clientPass))

	return err
}

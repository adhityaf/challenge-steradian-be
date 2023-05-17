package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(pass string) string {
	password := []byte(pass)

	hash, _ := bcrypt.GenerateFromPassword(password, 8)

	return string(hash)
}

func ComparePassword(hashed, password []byte) bool {
	h, p := []byte(hashed), []byte(password)

	err := bcrypt.CompareHashAndPassword(h, p)

	return err == nil
}

package middleware

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func PasswordHash(password string) string {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		log.Fatal(err)
	}

	return string(passwordHash)
}

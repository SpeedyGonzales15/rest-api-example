package test

import (
	"rest-api-example/pkg/middleware"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestPasswordHash(t *testing.T) {
	password := "test"
	hash := middleware.PasswordHash(password)

	if len(hash) == 0 {
		t.Error("empty hash")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		t.Error(err)
	}
}

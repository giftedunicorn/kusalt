package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (hash, salt, algorithm string, err error) {
	algorithm = "bcrypt"
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", "", "", err
	}

	return string(bytes), "", algorithm, nil
}

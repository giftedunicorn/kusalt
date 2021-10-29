package utils

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

// generate tokens
func TokenGenerator(userId string, email string, tokenType string) (tokenString string, err error) {
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var hmacSampleSecret []byte
	if tokenType == "refreshToken" {
		REFRESH_TOKEN_SECRET := os.Getenv("REFRESH_TOKEN_SECRET")
		hmacSampleSecret = []byte(REFRESH_TOKEN_SECRET)
	} else if tokenType == "accessToken" {
		ACCESS_TOKEN_SECRET := os.Getenv("ACCESS_TOKEN_SECRET")
		hmacSampleSecret = []byte(ACCESS_TOKEN_SECRET)
	}

	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":    userId,
		"email": email,
		"nbf":   time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err = token.SignedString(hmacSampleSecret)
	if err != nil {
		log.Fatal("Error generate token.")
	}

	return tokenString, err
}

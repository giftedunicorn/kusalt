package middlewares

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/giftedunicorn/kusalt/models"
	"github.com/giftedunicorn/kusalt/utils"
	"github.com/golang-jwt/jwt"
	"github.com/joho/godotenv"
)

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := godotenv.Load()
		if err != nil {
			log.Println("Error loading .env file")
			utils.RespondWithError(w, http.StatusInternalServerError, "")
			return
		}

		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			log.Println("Token not found.")
			utils.RespondWithError(w, http.StatusUnauthorized, "Token not found.")
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			ACCESS_TOKEN_SECRET := os.Getenv("ACCESS_TOKEN_SECRET")
			hmacSampleSecret := []byte(ACCESS_TOKEN_SECRET)
			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return hmacSampleSecret, nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			accessToken, err := models.GetAccessToken(tokenString)
			if err != nil {
				utils.RespondWithError(w, http.StatusUnauthorized, err.Error())
				return
			}
			if accessToken.Revoked == true {
				utils.RespondWithError(w, http.StatusUnauthorized, "Token revoked")
				return
			}

			ctx := r.Context()
			ctx = context.WithValue(ctx, "userId", claims["id"])
			r = r.WithContext(ctx)
		} else {
			utils.RespondWithError(w, http.StatusUnauthorized, "Authenticate failed")
			return
		}
		next.ServeHTTP(w, r)
	})
}

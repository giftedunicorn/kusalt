package middlewares

import (
	"crypto/md5"
	"encoding/hex"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/giftedunicorn/kusalt/utils"
	"github.com/joho/godotenv"
)

func getMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func reverseString(s string) string {
	rns := []rune(s) // convert to rune
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {

		// swap the letters of the string,
		// like first with last and so on.
		rns[i], rns[j] = rns[j], rns[i]
	}

	// return the reversed string.
	return string(rns)
}

func VerifySign(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := godotenv.Load()
		if err != nil {
			log.Println("Error loading .env file")
			utils.RespondWithError(w, http.StatusInternalServerError, "")
			return
		}

		sig := r.URL.Query().Get("sig")
		timestamp := r.URL.Query().Get("timestamp")
		method := r.Method

		if sig == "" {
			utils.RespondWithError(w, http.StatusUnauthorized, "Request not authorized")
			return
		}
		if timestamp == "" {
			utils.RespondWithError(w, http.StatusUnauthorized, "Request not authorized")
			return
		}

		API_KEY := os.Getenv("API_KEY")
		params := []string{"v1", "kusaltApp", timestamp, method, API_KEY}
		msg := strings.Join(params, "|")

		// reverse string
		reversed := reverseString(msg)

		msgSig := getMD5Hash(reversed)
		msgSig = getMD5Hash(msgSig)

		if sig != msgSig {
			utils.RespondWithError(w, http.StatusUnauthorized, "Request not authorized")
			return
		}

		next.ServeHTTP(w, r)
	})
}

package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"net/http"
)

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func GenerateSign(data interface{}) string {
	json, _ := json.Marshal(data)
	jsonStr := string(json)
	h := sha256.New()
	h.Write([]byte(jsonStr))
	return hex.EncodeToString(h.Sum(nil))
}

func RespondWithError(w http.ResponseWriter, code int, message string) {
	response := Response{
		Status:  "error",
		Data:    nil,
		Message: message,
	}

	respondWithJSON(w, code, response)
}

func RespondWithSuccess(w http.ResponseWriter, code int, data interface{}) {
	response := Response{
		Status:  "success",
		Data:    data,
		Message: "",
	}

	signature := GenerateSign(data)
	w.Header().Set("lv-signature", signature)

	respondWithJSON(w, code, response)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	payloadJson, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(payloadJson)
}

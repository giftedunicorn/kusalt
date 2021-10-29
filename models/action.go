package models

import (
	"log"
)

type Action struct {
	Id        string `json:"id"`
	UserId    string `json:"user_id"`
	ClientId  string `json:"client_id"`
	Platform  string `json:"platform"`
	Action    string `json:"action"`
	Value     string `json:"value"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func GetActions(userId string) ([]Action, error) {
	log.Println("GetActions invoked")

	return action, nil
}

func CreateAction(userId, userAction, value, clientId, platform string) error {
	log.Println("CreateAction invoked")

	return nil
}

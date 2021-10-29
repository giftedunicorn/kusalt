package models

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
)

type User struct {
	Id             string    `json:"id"`
	Email          string    `json:"email"`
	RefBy          MyNullInt `json:"ref_by"`
	Mfa            bool      `json:"mfa"`
	Verified       bool      `json:"verified"`
	Disabled       bool      `json:"disabled"`
	CreatedAt      string    `json:"created_at"`
	UpdatedAt      string    `json:"updated_at"`
}

func GetUser(email string, userId string) (*User, error) {
	log.Println("GetUser invoked")

}

func CreateUser(email string, refBy string, verified bool, referrer string) (*User, error) {
	log.Println("CreateUser invoked")

}

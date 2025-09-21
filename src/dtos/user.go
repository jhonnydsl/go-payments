package dtos

import "time"

type User struct {
	ID 			int `json:"id"`
	Email 		string `json:"email"`
	Name 		string `json:"name"`
	Password 	string `json:"password"`
	CreatedAt 	time.Time `json:"created_at"`
}

type UserInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserOutput struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	CreatedAt  time.Time `json:"created_at"`
}

type UserLogin struct {
	Email 		string `json:"email"`
	Password 	string `json:"password"`
}
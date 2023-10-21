package responses

import "github.com/starlingilcruz/golang-chat/internal/models"


type AuthResponse struct {
	User  models.User `json:"User"`
	Jwt 	string      `json:"Jwt"`
}
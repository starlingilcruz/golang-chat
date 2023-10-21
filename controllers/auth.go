package controllers

import (
	"net/http"
	"encoding/json"
	"fmt"


	"github.com/starlingilcruz/golang-chat/services"
	"github.com/starlingilcruz/golang-chat/utils"
	"github.com/starlingilcruz/golang-chat/http/requests"
	"github.com/starlingilcruz/golang-chat/http/responses"
)

type AuthController struct {
	// TODO dependency injection
	authService services.Auth
}

func (a *AuthController) RegisterService(s services.Auth) {
	a.authService = s
}

func (a *AuthController)SignUp(w http.ResponseWriter, r *http.Request) {
	var params requests.SignUpParams

	if err := utils.ParseBody(r, &params); err != nil {
		fmt.Println("Error parsing request payload")
		return
	}

	user, token, err := a.authService.SignUp(params.UserName, params.Email, params.Password)

	if err != nil {
		fmt.Printf("Has ocurrred an error while signing up user: %s", params.Email)
	}

	d, err := json.Marshal(responses.AuthResponse{User: user, Jwt: token})

	w.WriteHeader(http.StatusOK)
	w.Write(d)
}

func (a *AuthController)Login(w http.ResponseWriter, r *http.Request) {
	var params requests.LoginParams

	if err := utils.ParseBody(r, &params); err != nil {
		fmt.Println("Error parsing request payload")
		return
	}

	user, token, err := a.authService.Login(params.Email, params.Password)

	if err != nil {
		fmt.Printf("Has ocurrred an error while login user: %s", params.Email)
	}

	d, err := json.Marshal(responses.AuthResponse{User: user, Jwt: token})

	w.WriteHeader(http.StatusOK)
	w.Write(d)
}
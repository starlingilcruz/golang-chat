package services

import (
	"fmt"

	"github.com/starlingilcruz/golang-chat/internal/models"
	"github.com/starlingilcruz/golang-chat/utils"
)

type IAuth interface {
	Login(username, password string) (string, error)
	SignUp(username, email, password string) (string, error)
}

type Auth struct {
	repository   models.User
}

func (a *Auth) SignUp(username, email, password string) (models.User, string, error) {
	fmt.Println("Calling SignUp Service")
	
	// TODO support repo dependency injection

	var user models.User
	u := user.GetByEmail(email)

	if user.ID > 0 || u.Error != nil {
		fmt.Println("Signup: User already exists")
		return user, "", u.Error
	}

	newUser := models.User{
		UserName: username,
		Email:    email,
		Password: string(utils.HasPassword(password)),
	}

	if c := newUser.Create(); c.Error != nil {
		fmt.Println("Signup: Error creating new user")
		return newUser, "", u.Error
	}

	token, err := utils.GenerateJWT(newUser)

	if err != nil {
		fmt.Println("Error during token generation")
	}

	return newUser, token, u.Error
}

func (a *Auth) Login(email, password string) (models.User, string, error) {
	fmt.Println("Calling Login Service")

	var user models.User
	u := user.GetByEmail(email)

	if user.ID == 0 || u.Error != nil {
		fmt.Println("Login: Does not exists")
		return user, "", u.Error
	}

	if equal := utils.CompareHasAndPassword(user.Password, password); equal == false {
		fmt.Println("Login: Password is incorrect")
		return user, "", nil
	}

	token, err := utils.GenerateJWT(user)

	if err != nil {
		fmt.Println("Error during token generation")
	}

	return user, token, u.Error

}
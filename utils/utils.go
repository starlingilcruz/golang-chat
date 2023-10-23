package utils

import (
	"fmt"
	"encoding/json"

	"golang.org/x/crypto/bcrypt"
)

func HasPassword(pwd string) []byte {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)

	if err != nil {
		fmt.Println("Error while hashing password")
	}

	return hashPass
}

func CompareHasAndPassword(hashPwd string, pwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPwd), []byte(pwd))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		return false
	}
	return true
}

func ParseByteArray(r []byte, x interface{}) error {
	if err := json.Unmarshal(r, x); err != nil {
		return err
	}
	return nil
}
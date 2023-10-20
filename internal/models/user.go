package models

import (
	"gorm.io/gorm"
	// "gorm.io/driver/postgres"
)

type User struct {
	// https://gorm.io/docs/models.html#gorm-Model
	gorm.Model  // gorm defined model fields
	Username	string `json:"UserName,omitempty"`
	Email			string `json:"Email,omitempty"`
	Password	string `json:"Password,omitempty"`
}

func (u *User) GetUserByEmail(email string) *gorm.DB {

}
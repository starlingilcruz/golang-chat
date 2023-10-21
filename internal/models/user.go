package models

import (
	"gorm.io/gorm"

	"github.com/starlingilcruz/golang-chat/internal/db"
)

type User struct {
	// https://gorm.io/docs/models.html#gorm-Model
	gorm.Model  // gorm defined model fields
	UserName  string `json:"UserName,omitempty"`
	Email     string `json:"Email,omitempty"`
	Password  string `json:"Password,omitempty"`
}

func (u *User) Create() *gorm.DB {
	return db.GetInstance().Create(&u)
}

func (u *User) GetByUsername(username string) *gorm.DB {
	return db.GetInstance().Where("UserName = ?", username).First(&u)
}

func (u *User) GetByEmail(email string) *gorm.DB {
	r := db.GetInstance().Find(&u, User{Email: email})
	return r
	// return db.GetInstance().Where("Email = ?", email).First(&u)
}
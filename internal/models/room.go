package models

import (
	"gorm.io/gorm"

	"github.com/starlingilcruz/golang-chat/internal/db"

)

type Room struct {
	gorm.Model
	Name    string   `json:"Name,omitempty"`
	Members []User   `gorm:"many2many:room_users;"`
}

func (r *Room)Create() *gorm.DB {
	return db.GetInstance().Create(&r)
}

func (r *Room)List(rs *[]Room) *gorm.DB {
	return db.GetInstance().Find(&rs) 
}

func (r *Room)AddMember() *gorm.DB {
	res := db.GetInstance().Order("id DESC").Find(&r)
	return res 
} 
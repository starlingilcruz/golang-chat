package models

import (
	"gorm.io/gorm"

	"github.com/starlingilcruz/golang-chat/internal/db"

)

type Chat struct {
	gorm.Model
	Message    string   `json:"Message"`
	UserId     uint     `json:"UserId" gorm:"index"`
	User       User     `json:"User" gorm:"foreignKey:UserId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:NO ACTION;"`
	RoomId     uint     `json:"RoomId" gorm:"index"`
	Room       Room     `json:"Room" gorm:"foreignKey:RoomId;references:ID;constraint:OnUpdate:CASCADE,OnDelete:NO ACTION;"`
}

func (c *Chat)Create() *gorm.DB {
	return db.GetInstance().Create(&c)
}

func (c *Chat)List(roomId uint, ch *[]Chat) *gorm.DB {
	res := db.GetInstance().Where(Chat{RoomId: roomId}).Preload("Room").Preload("User").Order("id DESC").Find(&ch).Limit(50)
	return res 
}
package model

import "gorm.io/gorm"

type User struct {
	Name string `gorm:"varchar(15);unique;not null"`
	Password string `gorm:"varchar(20);not null"`
	Gender string `gorm:"default:'null'"`
	Email string `gorm:"type:varchar(20);unique index"`
	Age uint `gorm:"default:0"`
	Birthday int64 `gorm:"default:20000101"`
	Signature string `gorm:"type:varchar(100)"`
	HeadPhoto string `gorm:"type:varchar(100)"`
	Collect []Video `gorm:"many2many:user_collect"`
	Videos []Video
	Barrage []Barrage
	Comment []Comments
	Reply []Replies
	Black []User `gorm:"many2many:user_black"`
	gorm.Model
}

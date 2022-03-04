package model

import "gorm.io/gorm"

type Video struct {
	gorm.Model
	UserID uint
	Username string
	Title string `gorm:"varchar(50);not null"`
	VideoFile string `gorm:"varchar(100);not null"`
	Content string `gorm:"varchar(200)"`
	Click uint
	Like uint `gorm:"default:0"`
	Status bool `gorm:"default:false"`
	Comments []Comments
	Barrage []Barrage
}

package model

import "gorm.io/gorm"

type Comments struct {
	gorm.Model
	UserID uint
	VideoID uint
	Username string
	Content string `gorm:"varchar(100);not null"`
	Like uint `gorm:"default:0"`
	Dislike uint `gorm:"default:0"`
	Replies []Replies
}

type Replies struct {
	gorm.Model
	UserID uint
	CommentsID uint
	Username string
	Content string `gorm:"varchar(100);not null"`
	Like uint `gorm:"default:0"`
	Dislike uint `gorm:"default:0"`
}


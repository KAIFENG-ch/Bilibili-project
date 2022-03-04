package model

import "gorm.io/gorm"

type Barrage struct {
	Content string `gorm:"varchar(100);not null"`
	Time uint `gorm:"not null"`
	Color string `gorm:"type:varchar(10);not null"`
	UserID uint
	VideoID string
	gorm.Model
}

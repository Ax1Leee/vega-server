package model

import (
	"gorm.io/gorm"
)

type Review struct {
	gorm.Model
	UserID  uint
	User    User
	MovieID uint
	Movie   Movie
	Rating  float32
	Content string
}

func (Review) TableName() string {
	return "reviews"
}

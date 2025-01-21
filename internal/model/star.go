package model

import (
	"gorm.io/gorm"
)

type Star struct {
	gorm.Model
	Name       string
	MovieStars []MovieStar `gorm:"foreignKey:StarID"`
}

func (Star) TableName() string {
	return "stars"
}

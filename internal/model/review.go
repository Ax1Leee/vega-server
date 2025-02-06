package model

import (
	"gorm.io/gorm"
)

type Review struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt CustomTime
	UpdatedAt CustomTime
	DeletedAt gorm.DeletedAt `gorm:"index"`
	UserID    uint
	User      User
	MovieID   uint
	Movie     Movie
	Rating    float32
	Content   string
}

func (Review) TableName() string {
	return "reviews"
}

package model

import (
	"gorm.io/gorm"
)

type Star struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt CustomTime
	UpdatedAt CustomTime
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string
}

func (Star) TableName() string {
	return "stars"
}

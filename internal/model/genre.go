package model

import (
	"gorm.io/gorm"
)

type Genre struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt CustomTime
	UpdatedAt CustomTime
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string
}

func (Genre) TableName() string {
	return "genres"
}

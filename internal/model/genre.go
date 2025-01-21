package model

import (
	"gorm.io/gorm"
)

type Genre struct {
	gorm.Model
	Name        string
	MovieGenres []MovieGenre `gorm:"foreignKey:GenreID"`
}

func (Genre) TableName() string {
	return "genres"
}

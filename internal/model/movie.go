package model

import (
	"gorm.io/gorm"
)

type Movie struct {
	ID           uint `gorm:"primarykey"`
	CreatedAt    CustomTime
	UpdatedAt    CustomTime
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	Cover        string
	Title        string
	Genres       []Genre `gorm:"many2many:movie_genres"`
	ReleaseDate  string
	Location     string
	Director     string
	Stars        []Star `gorm:"many2many:movie_stars"`
	Language     string
	Runtime      string
	Storyline    string
	CriticRating float32
	UserRating   float32
	Reviews      []Review
}

func (Movie) TableName() string {
	return "movies"
}

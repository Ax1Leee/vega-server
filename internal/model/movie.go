package model

import (
	"gorm.io/gorm"
)

type Movie struct {
	gorm.Model
	Cover        string
	Title        string
	MovieGenres  []MovieGenre
	ReleaseDate  string
	Location     string
	Director     string
	MovieStars   []MovieStar
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

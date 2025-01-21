package model

import (
	"gorm.io/gorm"
	"time"
)

type MovieGenre struct {
	MovieID   uint  `gorm:"primaryKey;column:movie_id"`
	Movie     Movie `gorm:"foreignKey:MovieID"`
	GenreID   uint  `gorm:"primaryKey;column:genre_id"`
	Genre     Genre `gorm:"foreignKey:GenreID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (MovieGenre) TableName() string {
	return "movie_genres"
}

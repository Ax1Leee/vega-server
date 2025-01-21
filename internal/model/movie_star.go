package model

import (
	"gorm.io/gorm"
	"time"
)

type MovieStar struct {
	MovieID   uint  `gorm:"primaryKey;column:movie_id"`
	Movie     Movie `gorm:"foreignKey:MovieID"`
	StarID    uint  `gorm:"primaryKey;column:star_id"`
	Star      Star  `gorm:"foreignKey:StarID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (MovieStar) TableName() string {
	return "movie_stars"
}

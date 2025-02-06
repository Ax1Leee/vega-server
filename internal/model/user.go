package model

import (
	"gorm.io/gorm"
)

type User struct {
	ID             uint `gorm:"primarykey"`
	CreatedAt      CustomTime
	UpdatedAt      CustomTime
	DeletedAt      gorm.DeletedAt `gorm:"index"`
	Email          string
	HashedPassword string
	Avatar         string
	Name           string
	Gender         string
	BirthDate      string
	Location       string
	Bio            string
	Reviews        []Review
}

func (User) TableName() string {
	return "users"
}

package repository

import (
	"vega-server/internal/model"
)

type UserRepository struct {
	*Repository
}

func NewUserRepository(repository *Repository) *UserRepository {
	return &UserRepository{repository}
}

func (userRepository *UserRepository) Create(user *model.User) error {
	if err := userRepository.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (userRepository *UserRepository) Update(user *model.User) error {
	if err := userRepository.db.Save(user).Error; err != nil {
		return err
	}
	return nil
}

func (userRepository *UserRepository) Delete(user *model.User) error {
	if err := userRepository.db.Delete(user).Error; err != nil {
		return err
	}
	return nil
}

func (userRepository *UserRepository) QueryUserByID(id uint) (*model.User, error) {
	user := &model.User{}
	if err := userRepository.db.Preload("Reviews").Where("id = ?", id).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (userRepository *UserRepository) QueryUserByEmail(email string) (*model.User, error) {
	user := &model.User{}
	if err := userRepository.db.Where("email = ?", email).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

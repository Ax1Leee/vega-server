package repository

import (
	"vega-server/internal/model"
)

type ReviewRepository struct {
	*Repository
}

func NewReviewRepository(repository *Repository) *ReviewRepository {
	return &ReviewRepository{repository}
}

func (reviewRepository *ReviewRepository) Create(review *model.Review) error {
	if err := reviewRepository.db.Create(review).Error; err != nil {
		return err
	}
	return nil
}

func (reviewRepository *ReviewRepository) Update(review *model.Review) error {
	if err := reviewRepository.db.Save(review).Error; err != nil {
		return err
	}
	return nil
}

func (reviewRepository *ReviewRepository) Delete(review *model.Review) error {
	if err := reviewRepository.db.Delete(review).Error; err != nil {
		return err
	}
	return nil
}

func (reviewRepository *ReviewRepository) QueryReviewByID(id uint) (*model.Review, error) {
	review := &model.Review{}
	if err := reviewRepository.db.Preload("User, Movie").Where("id = ?", id).First(review).Error; err != nil {
		return nil, err
	}
	return review, nil
}

func (reviewRepository *ReviewRepository) QueryReviewByUserIDAndMovieID(userID uint, movieID uint) (*model.Review, error) {
	review := &model.Review{}
	if err := reviewRepository.db.Where("user_id = ? AND movie_id = ?", userID, movieID).First(review).Error; err != nil {
		return nil, err
	}
	return review, nil
}

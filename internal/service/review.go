package service

import (
	"gorm.io/gorm"
	"vega-server/api"
	"vega-server/internal/model"
	"vega-server/internal/repository"

	"errors"
)

type ReviewService struct {
	*Service
	reviewRepository *repository.ReviewRepository
}

func NewReviewService(service *Service, reviewRepository *repository.ReviewRepository) *ReviewService {
	return &ReviewService{
		Service:          service,
		reviewRepository: reviewRepository,
	}
}

func (reviewService *ReviewService) GetReviewFromUser(id uint) (*api.GetReviewFromUserResponseData, error) {
	review, err := reviewService.reviewRepository.QueryReviewByID(id)
	if err != nil {
		return nil, errors.New("failed to get review")
	}

	resp := &api.GetReviewFromUserResponseData{
		Movie: api.Movie{
			Cover: review.Movie.Cover,
			Title: review.Movie.Title,
		},
		Review: api.Review{
			Rating:    review.Rating,
			Content:   review.Content,
			UpdatedAt: review.UpdatedAt.Format("2025-01-01 01:01:01"),
		},
	}
	return resp, nil
}

func (reviewService *ReviewService) GetReviewToMovie(id uint) (*api.GetReviewToMovieResponseData, error) {
	review, err := reviewService.reviewRepository.QueryReviewByID(id)
	if err != nil {
		return nil, errors.New("failed to get review")
	}

	resp := &api.GetReviewToMovieResponseData{
		User: api.User{
			Avatar: review.User.Avatar,
			Name:   review.User.Name,
		},
		Review: api.Review{
			Rating:    review.Rating,
			Content:   review.Content,
			UpdatedAt: review.UpdatedAt.Format("2025-01-01 01:01:01"),
		},
	}
	return resp, nil
}

func (reviewService *ReviewService) GetReview(userID uint, movieID uint) (*api.GetReviewResponseData, error) {
	review, err := reviewService.reviewRepository.QueryReviewByUserIDAndMovieID(movieID, userID)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("failed to get review")
		} else {
			return nil, errors.New("review not found")
		}
	}

	resp := &api.GetReviewResponseData{
		Rating:  review.Rating,
		Content: review.Content,
	}
	return resp, nil
}

func (reviewService *ReviewService) SetReview(userID uint, movieID uint, req *api.SetReviewRequest) (*api.SetReviewResponseData, error) {
	review, err := reviewService.reviewRepository.QueryReviewByUserIDAndMovieID(userID, movieID)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("failed to set review")
		} else {
			review = &model.Review{
				UserID:  userID,
				MovieID: movieID,
				Rating:  req.Rating,
				Content: req.Content,
			}
			err = reviewService.reviewRepository.Create(review)
			if err != nil {
				return nil, errors.New("failed to set review")
			}

			resp := &api.SetReviewResponseData{
				Rating:  review.Rating,
				Content: review.Content,
			}
			return resp, nil
		}
	} else {
		review.Rating = req.Rating
		review.Content = req.Content
		err = reviewService.reviewRepository.Update(review)
		if err != nil {
			return nil, errors.New("failed to set review")
		}

		resp := &api.SetReviewResponseData{
			Rating:  review.Rating,
			Content: review.Content,
		}
		return resp, nil
	}
}

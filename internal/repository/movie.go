package repository

import (
	"vega-server/internal/model"

	"context"
	"fmt"
	"strconv"
)

type MovieRepository struct {
	*Repository
}

func NewMovieRepository(repository *Repository) *MovieRepository {
	return &MovieRepository{repository}
}

func (movieRepository *MovieRepository) QueryMovieByID(id uint) (*model.Movie, error) {
	movie := &model.Movie{}
	if err := movieRepository.db.Preload("Genres").Preload("Stars").Preload("Reviews").Where("id = ?", id).First(&movie).Error; err != nil {
		return nil, err
	}
	return movie, nil
}

func (movieRepository *MovieRepository) QueryIDsByGenreAndCategory(genre string, category string) ([]uint, error) {
	ids, err := movieRepository.rdb.LRange(context.Background(), fmt.Sprintf("genre:%s:category:%s", genre, category), 0, -1).Result()
	if err != nil {
		return nil, err
	}
	movieIDs := make([]uint, 0, len(ids))
	for _, id := range ids {
		movieID, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			return nil, err
		}
		movieIDs = append(movieIDs, uint(movieID))
	}
	return movieIDs, nil
}

func (movieRepository *MovieRepository) QueryTitleByID(id uint) (string, error) {
	title, err := movieRepository.rdb.Get(context.Background(), fmt.Sprintf("id:%d", id)).Result()
	if err != nil {
		return "", err
	}
	return title, nil
}

func (movieRepository *MovieRepository) QueryNowPlaying() ([]uint, error) {
	ids, err := movieRepository.rdb.LRange(context.Background(), "now-playing", 0, -1).Result()
	if err != nil {
		return nil, err
	}
	movieIDs := make([]uint, 0, len(ids))
	for _, id := range ids {
		movieID, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			return nil, err
		}
		movieIDs = append(movieIDs, uint(movieID))
	}
	return movieIDs, nil
}

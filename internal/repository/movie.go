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

func (movieRepository *MovieRepository) Create(movie *model.Movie) error {
	if err := movieRepository.db.Create(movie).Error; err != nil {
		return err
	}
	return nil
}

func (movieRepository *MovieRepository) Update(movie *model.Movie) error {
	if err := movieRepository.db.Save(movie).Error; err != nil {
		return err
	}
	return nil
}

func (movieRepository *MovieRepository) Delete(movie *model.Movie) error {
	if err := movieRepository.db.Delete(movie).Error; err != nil {
		return err
	}
	return nil
}

func (movieRepository *MovieRepository) QueryMovieByID(id uint) (*model.Movie, error) {
	movie := &model.Movie{}
	if err := movieRepository.db.Preload("Genres").Preload("Stars").Preload("Reviews").Where("id = ?", id).First(&movie).Error; err != nil {
		return nil, err
	}
	return movie, nil
}

func (movieRepository *MovieRepository) QueryMoviesByGenreAndCategory(genre string, category string) ([]uint, error) {
	ids, err := movieRepository.rdb.LRange(context.Background(), fmt.Sprintf("genre:%s:category:%s", genre, category), 0, -1).Result()
	if err != nil {
		return nil, err
	}
	movies := make([]uint, 0, len(ids))
	for _, id := range ids {
		movie, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			return nil, err
		}
		movies = append(movies, uint(movie))
	}
	return movies, nil
}

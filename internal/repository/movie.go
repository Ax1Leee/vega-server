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
	if err := movieRepository.db.Preload("MovieGenres, MovieGenres.Genre, MovieStars, MovieStars.Star, Reviews").Where("id = ?", id).First(&movie).Error; err != nil {
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
		movieID, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			return nil, err
		}
		movies = append(movies, uint(movieID))
	}
	return movies, nil
}

func (movieRepository *MovieRepository) QueryTitleByID(id uint) (string, error) {
	title, err := movieRepository.rdb.Get(context.Background(), fmt.Sprintf("id:%d", id)).Result()
	if err != nil {
		return "", err
	}
	return title, nil
}

func (movieRepository *MovieRepository) QueryHotMovies() ([]uint, error) {
	ids, err := movieRepository.rdb.LRange(context.Background(), "hot-movies", 0, -1).Result()
	if err != nil {
		return nil, err
	}
	movies := make([]uint, 0, len(ids))
	for _, id := range ids {
		movieID, err := strconv.ParseUint(id, 10, 64)
		if err != nil {
			return nil, err
		}
		movies = append(movies, uint(movieID))
	}
	return movies, nil
}

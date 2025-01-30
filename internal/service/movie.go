package service

import (
	"vega-server/api"
	"vega-server/internal/repository"

	"errors"
)

type MovieService struct {
	*Service
	movieRepository *repository.MovieRepository
}

func NewMovieService(service *Service, movieRepository *repository.MovieRepository) *MovieService {
	return &MovieService{
		Service:         service,
		movieRepository: movieRepository,
	}
}

func (movieService *MovieService) GetMovie(id uint) (*api.GetMovieResponseData, error) {
	movie, err := movieService.movieRepository.QueryMovieByID(id)
	if err != nil {
		return nil, errors.New("failed to get movie")
	}

	resp := &api.GetMovieResponseData{
		Movie: api.Movie{
			Cover:        movie.Cover,
			Title:        movie.Title,
			CriticRating: movie.CriticRating,
			UserRating:   movie.UserRating,
		},
	}
	return resp, nil
}

func (movieService *MovieService) GetAdvancedMovie(id uint) (*api.GetAdvancedMovieResponseData, error) {
	movie, err := movieService.movieRepository.QueryMovieByID(id)
	if err != nil {
		return nil, errors.New("failed to get movie")
	}

	genres := make([]string, 0, len(movie.MovieGenres))
	for _, movieGenre := range movie.MovieGenres {
		genres = append(genres, movieGenre.Genre.Name)
	}

	stars := make([]string, 0, len(movie.MovieStars))
	for _, movieStar := range movie.MovieStars {
		stars = append(stars, movieStar.Star.Name)
	}

	reviews := make([]uint, 0, len(movie.Reviews))
	for _, review := range movie.Reviews {
		reviews = append(reviews, review.ID)
	}

	resp := &api.GetAdvancedMovieResponseData{
		Movie: api.AdvancedMovie{
			Cover:        movie.Cover,
			Title:        movie.Title,
			Genres:       genres,
			ReleaseDate:  movie.ReleaseDate,
			Location:     movie.Location,
			Director:     movie.Director,
			Stars:        stars,
			Language:     movie.Language,
			Runtime:      movie.Runtime,
			Storyline:    movie.Storyline,
			CriticRating: movie.CriticRating,
			UserRating:   movie.UserRating,
		},
		Reviews: reviews,
	}
	return resp, nil
}

func (movieService *MovieService) GetIDs(genre string, category string) (*api.GetIDsResponseData, error) {
	ids, err := movieService.movieRepository.QueryIDsByGenreAndCategory(genre, category)
	if err != nil {
		return nil, errors.New("failed to get movies")
	}

	resp := &api.GetIDsResponseData{
		IDs: ids,
	}
	return resp, nil
}

func (movieService *MovieService) GetTitles(genre string, category string) (*api.GetTitlesResponseData, error) {
	ids, err := movieService.movieRepository.QueryIDsByGenreAndCategory(genre, category)
	if err != nil {
		return nil, errors.New("failed to get movies")
	}

	titles := make([]string, 0, 5)
	for _, id := range ids[0:5] {
		title, err := movieService.movieRepository.QueryTitleByID(id)
		if err != nil {
			return nil, errors.New("failed to get movies")
		}
		titles = append(titles, title)
	}

	resp := &api.GetTitlesResponseData{
		Titles: titles,
	}
	return resp, nil
}

func (movieService *MovieService) GetNowPlaying() (*api.GetNowPlayingResponseData, error) {
	movies, err := movieService.movieRepository.QueryNowPlaying()
	if err != nil {
		return nil, errors.New("failed to get movies")
	}

	resp := &api.GetNowPlayingResponseData{
		Movies: movies,
	}
	return resp, nil
}

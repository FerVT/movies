package usecase

import (
	"fmt"

	"github.com/FerVT/movies/model"

	"github.com/google/uuid"
)

type movies struct {
	moviesDB moviesDB
}

type moviesDB interface {
	GetMovieById(movieId string) (*model.Movie, error)
	GetAllMovies() ([]*model.Movie, error)
	CreateMovies(movies []*model.Movie) ([]*model.Movie, error)
	DeleteMoviesByIds(moviesIds []string) error
}

func NewMovies(mDB moviesDB) *movies {
	return &movies{moviesDB: mDB}
}

func (u *movies) GetAllMovies() ([]*model.Movie, error) {
	movies, err := u.moviesDB.GetAllMovies()
	if err != nil {
		return nil, fmt.Errorf("usecase.GetAllMovies(): %w", err)
	}

	return movies, nil
}

func (u *movies) GetMovie(movieId string) (*model.Movie, error) {
	movie, err := u.moviesDB.GetMovieById(movieId)
	if err != nil {
		return nil, fmt.Errorf("usecase.GetMovie(): %w", err)
	}

	return movie, nil
}

func (u *movies) CreateMovies(movies []*model.Movie) ([]*model.Movie, error) {
	for _, m := range movies {
		m.ID = uuid.NewString()
	}

	movies, err := u.moviesDB.CreateMovies(movies)
	if err != nil {
		return nil, fmt.Errorf("usecase.CreateMovies(): %w", err)
	}

	return movies, nil
}

func (u *movies) DeleteMovies(moviesIds []string) error {
	err := u.moviesDB.DeleteMoviesByIds(moviesIds)
	if err != nil {
		return fmt.Errorf("usecase.DeleteMovies(): %w", err)
	}

	return nil
}

package usecase

import (
	"github.com/FerVT/movies/model"

	"github.com/google/uuid"
)

type movies struct {
	moviesDB moviesDB
}

type moviesDB interface {
	GetMovieByID(movieId string) (*model.Movie, error)
	GetAllMovies() ([]*model.Movie, error)
	CreateMovies(movies []*model.Movie) ([]*model.Movie, error)
	DeleteMoviesByIds(moviesIds []string) error
}

func NewMovies(mDB moviesDB) *movies {
	return &movies{moviesDB: mDB}
}

func (u *movies) GetAllMovies() ([]*model.Movie, error) {
	return u.moviesDB.GetAllMovies()
}

func (u *movies) GetMovie(movieId string) (*model.Movie, error) {
	return u.moviesDB.GetMovieByID(movieId)
}

func (u *movies) CreateMovies(movies []*model.Movie) ([]*model.Movie, error) {
	for _, m := range movies {
		m.ID = uuid.NewString()
	}

	return u.moviesDB.CreateMovies(movies)
}

func (u *movies) DeleteMovies(moviesIds []string) error {
	return u.moviesDB.DeleteMoviesByIds(moviesIds)
}

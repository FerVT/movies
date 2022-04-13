package usecase

import "github.com/FerVT/movies/model"

type Movies struct {
	moviesDB moviesDB
}

type moviesDB interface {
	GetMovieByID(movieId string) (*model.Movie, error)
	GetAllMovies() ([]*model.Movie, error)
}

func NewMovies(mDB moviesDB) *Movies {
	return &Movies{moviesDB: mDB}
}

func (u *Movies) GetAllMovies() ([]*model.Movie, error) {
	return u.moviesDB.GetAllMovies()
}

func (u *Movies) GetMovie(movieId string) (*model.Movie, error) {
	return u.moviesDB.GetMovieByID(movieId)
}

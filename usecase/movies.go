package usecase

import "github.com/FerVT/movies/model"

type Movies struct{}

func NewMovies() *Movies {
	return &Movies{}
}

func (u *Movies) GetAllMovies() ([]*model.Movie, error) {
	return nil, nil
}

func (u *Movies) GetMovie(movieId string) (*model.Movie, error) {
	return nil, nil
}

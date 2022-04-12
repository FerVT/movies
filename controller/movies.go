package controller

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/FerVT/movies/model"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/unrolled/render"
)

type Movies struct {
	render  *render.Render
	usecase moviesUsecase
}

type moviesUsecase interface {
	GetAllMovies() ([]*model.Movie, error)
	GetMovie(movieId string) (*model.Movie, error)
}

func NewMovies(r *render.Render, u moviesUsecase) *Movies {
	return &Movies{
		render:  r,
		usecase: u,
	}
}

func (c *Movies) GetAllMovies(w http.ResponseWriter, req *http.Request) {
	log.Info("get all movies controller")

	movies, err := c.usecase.GetAllMovies()
	if err != nil {
		log.Error(err)

		c.render.Text(w, http.StatusInternalServerError, "unexpected error fetching movies")
		return
	}

	if len(movies) == 0 {
		log.Info("movies not found")

		c.render.Text(w, http.StatusNotFound, "movies not found")
		return
	}

	c.render.JSON(w, http.StatusOK, movies)
}

func (c *Movies) GetMovie(w http.ResponseWriter, req *http.Request) {
	log.Info("get movie controller")

	movieID := mux.Vars(req)["id"]
	if strings.TrimSpace(movieID) == "" {
		log.Info("invalid movieID param")

		c.render.Text(w, http.StatusBadRequest, "invalid id param")
		return
	}

	movie, err := c.usecase.GetMovie(movieID)
	if err != nil {
		log.Error(err)

		c.render.Text(w, http.StatusInternalServerError, "unexpected error fetching movie")
		return
	}

	if movie == nil {
		log.Info(fmt.Sprintf("movie not found: %s", movieID))

		c.render.Text(w, http.StatusNotFound, "movie not found")
		return
	}

	c.render.JSON(w, http.StatusOK, movie)
}

func (c *Movies) CreateMovies(w http.ResponseWriter, req *http.Request) {
	log.Info("create movies controller")

	c.render.Text(w, http.StatusCreated, "Movies created fam")
}

func (c *Movies) DeleteMovies(w http.ResponseWriter, req *http.Request) {
	log.Info("delete movies controller")

	c.render.Text(w, http.StatusNoContent, "")
}

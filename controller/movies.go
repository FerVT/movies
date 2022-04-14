package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/FerVT/movies/model"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/unrolled/render"
)

type movies struct {
	render  *render.Render
	usecase moviesUsecase
}

type moviesUsecase interface {
	GetAllMovies() ([]*model.Movie, error)
	GetMovie(movieId string) (*model.Movie, error)
	CreateMovies(movies []*model.Movie) ([]*model.Movie, error)
	DeleteMovies(moviesIds []string) error
}

func NewMovies(r *render.Render, u moviesUsecase) *movies {
	return &movies{
		render:  r,
		usecase: u,
	}
}

func (c *movies) GetAllMovies(w http.ResponseWriter, req *http.Request) {
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

func (c *movies) GetMovie(w http.ResponseWriter, req *http.Request) {
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

func (c *movies) CreateMovies(w http.ResponseWriter, req *http.Request) {
	log.Info("create movies controller")

	var movies []*model.Movie
	err := json.NewDecoder(req.Body).Decode(&movies)
	if err != nil {
		log.Error(err)

		c.render.Text(w, http.StatusBadRequest, "invalid request body")
		return
	}

	for _, m := range movies {
		err = m.ValidateFields()
		if err != nil {
			log.Error(err)

			c.render.Text(w, http.StatusBadRequest, "invalid request body: "+err.Error())
			return
		}
	}

	movies, err = c.usecase.CreateMovies(movies)
	if err != nil {
		log.Error(err)

		c.render.Text(w, http.StatusInternalServerError, "unexpected error creating movies")
		return
	}

	c.render.JSON(w, http.StatusCreated, movies)
}

func (c *movies) DeleteMovies(w http.ResponseWriter, req *http.Request) {
	log.Info("delete movies controller")

	ids := strings.Split(mux.Vars(req)["ids"], ",")
	if len(ids) == 0 {
		log.Info("empty ids list to delete")

		c.render.Text(w, http.StatusNoContent, "")
		return
	}

	for _, id := range ids {
		if strings.TrimSpace(id) == "" {
			log.Info("empty id in ids list")

			c.render.Text(w, http.StatusBadRequest, "empty id in ids list")
			return
		}
	}

	err := c.usecase.DeleteMovies(ids)
	if err != nil {
		log.Error(err)

		c.render.Text(w, http.StatusInternalServerError, "unexpected error deleting movies")
		return
	}

	c.render.Text(w, http.StatusNoContent, "")
}

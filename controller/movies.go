package controller

import (
	"net/http"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/unrolled/render"
)

type Movies struct {
	render *render.Render
}

func NewMovies(render *render.Render) *Movies {
	return &Movies{
		render: render,
	}
}

func (c *Movies) GetAllMovies(w http.ResponseWriter, req *http.Request) {
	log.Info("get all movies controller")

	c.render.Text(w, http.StatusOK, "Here are your movies m8")
}

func (c *Movies) GetMovie(w http.ResponseWriter, req *http.Request) {
	log.Info("get movie controller")

	movieID := mux.Vars(req)["id"]

	c.render.Text(w, http.StatusOK, "Here is your movie m8: "+movieID)
}

func (c *Movies) CreateMovies(w http.ResponseWriter, req *http.Request) {
	log.Info("create movies controller")

	c.render.Text(w, http.StatusCreated, "Movies created fam")
}

func (c *Movies) UpdateMovies(w http.ResponseWriter, req *http.Request) {
	log.Info("update movies controller")

	c.render.Text(w, http.StatusCreated, "Movies updated fam")
}

func (c *Movies) DeleteMovies(w http.ResponseWriter, req *http.Request) {
	log.Info("delete movies controller")

	c.render.Text(w, http.StatusNoContent, "")
}

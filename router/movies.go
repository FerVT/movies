package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

type socksController interface {
	GetAllMovies(w http.ResponseWriter, req *http.Request)
	GetMovie(w http.ResponseWriter, req *http.Request)
	CreateMovies(w http.ResponseWriter, req *http.Request)
	UpdateMovies(w http.ResponseWriter, req *http.Request)
	DeleteMovies(w http.ResponseWriter, req *http.Request)
}

func GetRouter(sc socksController) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/movies", sc.GetAllMovies).Methods(http.MethodGet).Name("GetAllMovies")
	r.HandleFunc("/movies/{id}", sc.GetMovie).Methods(http.MethodGet).Name("GetMovie")
	r.HandleFunc("/movies", sc.CreateMovies).Methods(http.MethodPost).Name("CreateMovies")
	r.HandleFunc("/movies", sc.UpdateMovies).Methods(http.MethodPut).Name("UpdateMovies")
	r.HandleFunc("/movies", sc.DeleteMovies).Methods(http.MethodDelete).Name("DeleteMovies")

	return r
}

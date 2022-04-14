package main

import (
	"net/http"

	"github.com/FerVT/movies/config"
	"github.com/FerVT/movies/controller"
	"github.com/FerVT/movies/router"
	"github.com/FerVT/movies/service/ddb"
	"github.com/FerVT/movies/usecase"

	log "github.com/sirupsen/logrus"
	"github.com/unrolled/render"
)

func main() {
	renderer := render.New()

	appConfig, err := config.GetConfig("config.json")
	if err != nil {
		log.Fatal("main(): ", err)
		return
	}

	moviesDB, err := ddb.NewMovies(appConfig)
	if err != nil {
		log.Fatal("main(): ", err)
		return
	}

	moviesUsecase := usecase.NewMovies(moviesDB)

	moviesController := controller.NewMovies(renderer, moviesUsecase)

	appRouter := router.GetRouter(moviesController)

	log.Info("starting API server on port " + appConfig.AppPort)
	log.Fatal(http.ListenAndServe(":"+appConfig.AppPort, appRouter))
}

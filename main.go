package main

import (
	"net/http"

	log "github.com/sirupsen/logrus"
	"github.com/unrolled/render"

	"github.com/FerVT/movies/config"
	"github.com/FerVT/movies/controller"
	"github.com/FerVT/movies/router"
)

func main() {
	renderer := render.New()

	appConfig, err := config.GetConfig("config.json")
	if err != nil {
		log.Fatal("error loading configuration: ", err)
		return
	}

	moviesController := controller.NewMovies(renderer)

	appRouter := router.GetRouter(moviesController)

	log.Info("starting API server on port " + appConfig.AppPort)
	log.Fatal(http.ListenAndServe(":"+appConfig.AppPort, appRouter))
}

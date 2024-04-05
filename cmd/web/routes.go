package main

import (
	"dshusdock/tw_prac1/config"
	"dshusdock/tw_prac1/internal/handlers"
	"dshusdock/tw_prac1/internal/handlers/layoutvw"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func initRouteHandlers() {
	app.ViewCache["lyoutvw"] = layoutvw.AppLayoutVw.RegisterView(app)
}


func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Get("/", handlers.Repo.Home)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}

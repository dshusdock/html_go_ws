package main

import (
	"dshusdock/tw_prac1/config"
	"dshusdock/tw_prac1/internal/constants"
	"dshusdock/tw_prac1/internal/handlers"
	"fmt"
	"log"
	"log/slog"
	"net/http"
)

// AppConfig holds the application config

const portNumber = ":8082"
var app config.AppConfig 

func main()  {
	app.InProduction = false
	app.ViewCache = make(map[string]constants.ViewInteface)

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	var programLevel = new(slog.LevelVar) // Info by default
	programLevel.Set(slog.LevelDebug)

	initRouteHandlers()
	
	fmt.Printf("Staring application on port %s\n", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
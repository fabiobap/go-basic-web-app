package main

import (
	"fmt"
	"log"
	"myapp/config"
	"myapp/pkg/handlers"
	"myapp/pkg/render"
	"net/http"
)

const HTTP_PORT = ":8080"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)

	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

	fmt.Printf("Starting application on port: %s", HTTP_PORT)

	srv := &http.Server{
		Addr:    HTTP_PORT,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}

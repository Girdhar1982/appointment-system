package main

import (
	"net/http"
	"github.com/girdhar1982/appointment-system/config"
	"github.com/girdhar1982/appointment-system/internal/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler {
mux := chi.NewRouter()
mux.Use(middleware.Recoverer)
mux.Use(NoSurf)
mux.Use(SessionLoad)
mux.Get("/",handlers.Repo.Home)
mux.Get("/about",handlers.Repo.About)
mux.Get("/contact",handlers.Repo.Contact)
mux.Get("/generals",handlers.Repo.Generals)
mux.Get("/majors",handlers.Repo.Majors)
mux.Get("/make-reservation",handlers.Repo.Reservations)
mux.Get("/search-availability",handlers.Repo.Availability)
mux.Post("/search-availability",handlers.Repo.PostAvailability)
mux.Post("/search-availability-json",handlers.Repo.JsonAvailability)
fileServer := http.FileServer(http.Dir("./static/"))
mux.Handle("/static/*",http.StripPrefix("/static",fileServer));
return mux;
}
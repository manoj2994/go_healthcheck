package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/manoj2994/go_healthcheck/db"
	"github.com/manoj2994/go_healthcheck/handlers"
)

func main() {
	DB, _ := db.Connect()
	h := &handlers.Handler{DB: DB}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/healthz", h.Healthcheckwrite)

	r.Get("/deepz", h.Healthcheckping)

	http.ListenAndServe(":8080", r)
}

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/manoj2994/go_healthcheck/db"
	"github.com/manoj2994/go_healthcheck/handlers"
)

func main() {
	DB, _ := db.Connect()

	err := DB.Ping()
	if err != nil {
		log.Fatalf("Failed to ping DB: %v", err)
	} else {
		fmt.Println("Successfully pinged the database!")
	}
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/healthz", handlers.Healthcheck)

	http.ListenAndServe(":8080", r)
}

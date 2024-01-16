package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func main() {
	r := chi.NewRouter();
	err := godotenv.Load();

	if err != nil {
		log.Fatal("Could not load .env file");
	}

	r.Use(middleware.Logger)
	r.Use(middleware.StripSlashes)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("SERVICE OK"));
	});

	http.ListenAndServe(os.Getenv("PORT"), r);
}
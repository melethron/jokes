package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/melethron/jokes.git/internal/api/jokes"
	"github.com/melethron/jokes.git/internal/config"
	"github.com/melethron/jokes.git/internal/handler"
)

func main() {
	cfg := config.Server{}
	err := cleanenv.ReadConfig("config.yml", &cfg)
	if err != nil {
		log.Fatal(err)
	}
	r := chi.NewRouter()
	apiClient := jokes.NewJokeClient(cfg.JokeUrl)
	h := handler.NewHandler(apiClient)
	r.Get("/hello", h.Hello)
	path := cfg.Host + ":" + cfg.Port
	log.Printf("Starting server at %s", path)
	if err := http.ListenAndServe(path, r); err != nil {
		log.Fatal(err)
	}
	log.Print("Shutting down server")
}

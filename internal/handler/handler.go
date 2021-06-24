package handler

import (
	"fmt"
	"net/http"

	"github.com/melethron/jokes.git/internal/api"
)

type Handler struct {
	JokeClient api.Client
}

func NewHandler(jc api.Client) *Handler {
	return &Handler{
		JokeClient: jc,
	}
}

func (h *Handler) Hello(w http.ResponseWriter, r *http.Request) {
	joke, err := h.JokeClient.GetJoke()
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	fmt.Fprint(w, joke.Joke)
}

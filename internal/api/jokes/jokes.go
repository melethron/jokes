package jokes

import (
	"encoding/json"
	"net/http"

	"github.com/melethron/jokes.git/internal/api"
)

type JokeClient struct {
	url string
}

const getJokePath = "/api?format=json"

func NewJokeClient(baseUrl string) *JokeClient {
	return &JokeClient{
		url: baseUrl,
	}
}

func (jc *JokeClient) GetJoke() (*api.JokeResponse, error) {
	urlPath := jc.url + getJokePath
	resp, err := http.Get(urlPath)
	if err != nil {
		return nil, err
	}
	var data api.JokeResponse

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

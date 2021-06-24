package config

type Server struct {
	Port    string `yaml:"port" env:"PORT"`
	Host    string `yaml:"host" env:"HOST" env-default:"127.0.0.1"`
	JokeUrl string `yaml:"jokeURL" env:"JOKE_URL"`
}

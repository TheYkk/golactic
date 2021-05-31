package config

type Http struct {
	Port string `env:"PORT,default=8080"`
}

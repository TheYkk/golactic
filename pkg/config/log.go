package config

type Logger struct {
	Level string `env:"LEVEL,default=info"`
}

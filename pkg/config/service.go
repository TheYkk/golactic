package config

type Service struct {
	Name    string
	Version string
	Project string `env:"PROJECT,default=github.com/theykk/galactic"`
}

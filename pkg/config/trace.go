package config

type Trace struct {
	URl       string `env:"URL"`
	BatchSize int `env:"BATCH_SIZE,default=16"`
}

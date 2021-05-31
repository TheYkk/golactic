package config

import (
	"context"
	"github.com/rs/zerolog/log"
	"github.com/sethvargo/go-envconfig"
)

type Configuration struct {
	Tracer  Trace
	Log     Logger
	Service Service
	Http    Http
}

func InitConfig() Configuration {
	ctx := context.Background()

	var c Configuration
	if err := envconfig.Process(ctx, &c); err != nil {
		log.Fatal().Err(err)
	}

	return c
}

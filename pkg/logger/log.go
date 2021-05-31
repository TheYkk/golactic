package logger

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/theykk/golactic/pkg/config"
)

func InitLogger(cnf *config.Configuration) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs
	level, err := zerolog.ParseLevel(cnf.Log.Level)
	if err != nil {
		log.Fatal().Err(err)
	}
	zerolog.SetGlobalLevel(level)
}

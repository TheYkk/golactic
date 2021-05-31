package main

import (
	"flag"
	"github.com/rs/zerolog/log"
	"github.com/theykk/golactic/pkg/config"
	"github.com/theykk/golactic/pkg/logger"
	"github.com/theykk/golactic/pkg/metric"
	"github.com/theykk/golactic/pkg/tracing"
	"net/http"
	"os"
)

var (
	Version     = "Dev"
	ServiceName = "ticker"
)

var printVersion = flag.Bool("v", false, "Print version")
var help = flag.Bool("help", false, "Get Help")

func init() {
	flag.Usage = func() {
		flag.PrintDefaults()
		os.Exit(0)
	}

	flag.Parse()

	if *help {
		flag.PrintDefaults()
		os.Exit(0)
	}

	if *printVersion {
		println(Version)
		os.Exit(0)
	}

}

func main() {
	// Init config
	configuration := config.InitConfig()
	configuration.Service.Name = ServiceName
	configuration.Service.Version = Version

	// Init Logger
	logger.InitLogger(&configuration)

	//Init Tracing
	tracing.InitTracer(&configuration)

	// Init Metric
	metric.InitMetric()
	// Init GRPC

	// Init HTTP layer for metrics and health check
	if err := http.ListenAndServe(":"+configuration.Http.Port, nil); err != nil {
		log.Fatal().Err(err)
	}
	// Gracefull shutdown
}

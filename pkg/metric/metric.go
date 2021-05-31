package metric

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

var (
	reqCounterHttp = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Count of all HTTP requests.",
	}, []string{"code", "method"})

	reqCounterGRPC = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "grpc_requests_total",
		Help: "Count of all GRPC requests.",
	}, []string{"code", "method"})

	reqDurations = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Name: "requests_durations",
			Help: "Requests latencies in seconds",
		})

	authnDurations = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Name: "authentication_durations",
			Help: "Requests authentication latencies in seconds",
		})

	authzDurations = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Name: "authorization_durations",
			Help: "Requests authorization latencies in seconds",
		})
)

func InitMetric() {
	prometheus.MustRegister(reqCounterHttp)
	prometheus.MustRegister(reqDurations)
	prometheus.MustRegister(reqCounterGRPC)
	prometheus.MustRegister(authnDurations)
	prometheus.MustRegister(authzDurations)
	http.Handle("/metrics", promhttp.Handler())
}

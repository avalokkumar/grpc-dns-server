package metrics

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	DNSQueryDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "dns_query_duration_seconds",
			Help:    "Histogram of DNS query durations.",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"type"},
	)
)

func InitMetrics() {
	prometheus.MustRegister(DNSQueryDuration)
	http.Handle("/metrics", promhttp.Handler())
	go http.ListenAndServe(":9090", nil)
}
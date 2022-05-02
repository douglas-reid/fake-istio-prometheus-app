package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	istioStdLabels = []string{
		"connection_security_policy",
		"destination_app",
		"destination_canonical_service",
		"destination_canonical_revision",
		"destination_principal",
		"destination_service",
		"destination_service_name",
		"destination_service_namespace",
		"destination_version",
		"destination_workload",
		"destination_workload_namespace",
		"reporter",
		"request_protocol",
		"response_code",
		"response_flags",
		"source_app",
		"source_canonical_service",
		"source_canonical_revision",
		"source_principal",
		"source_version",
		"source_workload",
		"source_workload_namespace",
	}

	defaultValues = []string{
		"mutual_tls",
		"details",
		"details",
		"v1",
		"cluster.local/ns/default/sa/default",
		"details.default.svc.cluster.local",
		"details",
		"default",
		"v1",
		"details-v1",
		"default",
		"destination",
		"http",
		"200",
		"-",
		"productpage",
		"productpage",
		"v1",
		"cluster.local/ns/default/sa/default",
		"v1",
		"productpage-v1",
		"default",
	}

	reqs = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "istio_requests_total",
	}, istioStdLabels)

	dur = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name: "istio_request_duration_milliseconds",
	}, istioStdLabels)
)

func generateBaseline() {
	reqs.WithLabelValues(defaultValues...).Inc()
	dur.WithLabelValues(defaultValues...).Observe(rand.Float64() * 10)
}

func main() {
	go func() {
		for {
			generateBaseline()
			time.Sleep(500 * time.Millisecond)
		}
	}()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":15000", nil)
}

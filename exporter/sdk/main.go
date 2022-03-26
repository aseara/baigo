package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func main() {
	reg := prometheus.NewRegistry()

	queueLength := prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "app",
		Subsystem: "cube_demo",
		Name:      "queue_length",
		Help:      "The number of items in the queue.",
		ConstLabels: map[string]string{
			"module": "http-server",
		},
	})
	_ = reg.Register(queueLength)

	queueLength.Set(100)

	totalRequest := prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: "app",
		Subsystem: "cube_demo",
		Name:      "http_requests_total",
		Help:      "The total number of handled HTTP requests.",
		ConstLabels: map[string]string{
			"module": "http-server",
		},
	})
	_ = reg.Register(totalRequest)
	totalRequest.Add(100)

	requestDurations := prometheus.NewHistogram(prometheus.HistogramOpts{
		Namespace: "app",
		Subsystem: "cube_demo",
		Name:      "http_requests_duration_seconds",
		Help:      "A histogram of the HTTP request durations in seconds.",
		Buckets:   []float64{0.05, 0.1, 0.25, 0.5, 1, 2.5, 5, 10},
	})
	_ = reg.Register(requestDurations)
	requestDurations.Observe(0.1)
	requestDurations.Observe(0.2)
	requestDurations.Observe(3)
	requestDurations.Observe(9)

	requestDurations2 := prometheus.NewSummary(prometheus.SummaryOpts{
		Namespace: "app",
		Subsystem: "cube_demo",
		Name:      "http_requests_duration_seconds2",
		Help:      "A summary of the HTTP request durations in seconds.",
		Objectives: map[float64]float64{
			0.5:  0.05,
			0.9:  0.01,
			0.99: 0.001,
		},
	})

	for _, v := range []float64{0.01, 0.02, 0.3, 0.4, 0.6, 0.7, 5.5, 11} {
		requestDurations2.Observe(v)
	}

	_ = reg.Register(requestDurations2)

	reg.Register(collectors.NewGoCollector())

	http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{Registry: reg}))
	_ = http.ListenAndServe(":8050", nil)
}

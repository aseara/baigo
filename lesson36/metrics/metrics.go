package metrics

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	// ClientConnected client connected gauge
	ClientConnected prometheus.Gauge
	// ReqRecvTotal   req rec counter
	ReqRecvTotal prometheus.Counter
	// RspSendTotal rsp send counter
	RspSendTotal prometheus.Counter
)

var metricsHTTPPort int32 = 8889

func init() {
	ReqRecvTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "tcp_server_demo2_req_recv_total",
	})
	RspSendTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "tcp_server_demo2_rsp_send_total",
	})
	ClientConnected = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "tcp_server_demo2_client_connected",
	})

	prometheus.MustRegister(ReqRecvTotal, RspSendTotal, ClientConnected)

	mu := http.NewServeMux()
	mu.Handle("/metrics", promhttp.Handler())

	metricsServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", metricsHTTPPort),
		Handler: mu,
	}

	go func() {
		err := metricsServer.ListenAndServe()
		if err != nil {
			fmt.Println("prometheus-exporter http server start failed:", err)
		}
	}()
	fmt.Println("metrics server start ok(*:8889)")
}

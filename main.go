package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Labels map[string]string

var (
	httpRequestsCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "http_requests",
		Help: "number of http requests",
	})
)

func init() {
	// Metrics have to be registered to be exposed:
	prometheus.MustRegister(httpRequestsCounter)
}

func main() {
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		defer httpRequestsCounter.Inc()
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}

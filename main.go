package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	port := flag.String("port", "2113", "prometheus port")
	flag.Parse()
	//Create a new instance of the foocollector and
	//register it with the prometheus client.
	foo := newFooCollector()
	prometheus.MustRegister(foo)

	//This section will start the HTTP server and expose
	//any metrics on the /metrics endpoint.
	http.Handle("/metrics", promhttp.Handler())
	fmt.Println("Beginning to serve on port " + *port)
	fmt.Println(http.ListenAndServe(":"+*port, nil))
}

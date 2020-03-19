package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
)

//Define a struct for you collector that contains pointers
//to prometheus descriptors for each metric you wish to expose.
//Note you can also include fields of other types if they provide utility
//but we just won't be exposing them as metrics.
type FooCollector struct {
	ylowMetric  *prometheus.Desc
	yavgMetric  *prometheus.Desc
	yhighMetric *prometheus.Desc
	barMetric   *prometheus.Desc
}

//You must create a constructor for you collector that
//initializes every descriptor and returns a pointer to the collector
func newFooCollector() *FooCollector {
	return &FooCollector{
		ylowMetric: prometheus.NewDesc("ylow_metric",
			"Shows whether a ylow has occurred in our cluster",
			nil, nil,
		),
		yavgMetric: prometheus.NewDesc("yavg_metric",
			"Shows whether a yavg has occurred in our cluster",
			nil, nil,
		),
		yhighMetric: prometheus.NewDesc("yhigh_metric",
			"Shows whether a yhigh has occurred in our cluster",
			nil, nil,
		),
		barMetric: prometheus.NewDesc("bar_metric",
			"Shows whether a bar has occurred in our cluster",
			nil, nil,
		),
	}
}

func getLine() string {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), "frame") {
			return scanner.Text()
		}
	}
	return "empty"
}

func convertLine(Line string) []float64 {
	var numbers []float64
	lst := strings.Split(Line, ",")[30:]
	for _, metric := range lst {
		if n, err := strconv.ParseFloat(metric, 64); err == nil {
			numbers = append(numbers, n)
		}
	}
	return numbers
}

//Each and every collector must implement the Describe function.
//It essentially writes all descriptors to the prometheus desc channel.
func (collector *FooCollector) Describe(ch chan<- *prometheus.Desc) {

	//Update this section with the each metric you create for a given collector
	ch <- collector.ylowMetric
	ch <- collector.yavgMetric
	ch <- collector.yhighMetric
	ch <- collector.barMetric
}

//Collect implements required collect function for all promehteus collectors
func (collector *FooCollector) Collect(ch chan<- prometheus.Metric) {

	//Implement logic here to determine proper metric value to return to prometheus
	//for each descriptor or call other functions that do so.
	//Write latest value for each metric in the prometheus metric channel.
	//Note that you can pass CounterValue, GaugeValue, or UntypedValue types here.
	line := getLine()
	if line != "empty" {
		metrics := convertLine(line)
		ch <- prometheus.MustNewConstMetric(collector.ylowMetric, prometheus.CounterValue, metrics[30])
		ch <- prometheus.MustNewConstMetric(collector.yavgMetric, prometheus.CounterValue, metrics[31])
		ch <- prometheus.MustNewConstMetric(collector.yhighMetric, prometheus.CounterValue, metrics[32])
		ch <- prometheus.MustNewConstMetric(collector.barMetric, prometheus.CounterValue, metrics[32])
	} else {
		log.Printf("line: %v", line)
	}

}

// 22-ymin

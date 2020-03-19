package main

import (
	"bufio"
	"fmt"
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
func getYlow() float64 {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), "frame") {
			Val, _ := strconv.ParseFloat(strings.Split(scanner.Text(), ",")[30], 64) //[29], 64)
			//fmt.Print(Val)
			return Val
		}
	}
	return 0.5
}
func getYavg() float64 {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), "frame") {
			Val, _ := strconv.ParseFloat(strings.Split(scanner.Text(), ",")[31], 64) //[29], 64)
			//fmt.Print(Val)
			return Val
		}
	}
	return 0.5
}
func getYhigh() float64 {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if strings.HasPrefix(scanner.Text(), "frame") {
			Val, _ := strconv.ParseFloat(strings.Split(scanner.Text(), ",")[32], 64) //[29], 64)
			//fmt.Print(Val)
			return Val
		}
	}
	return 0.5
}

func getVals() []float64 {
	var metrics []float64
	var cur float64
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	if strings.HasPrefix(line, "frame") {
		s := strings.Split(line, ",")[29:]
		//println("blaaa")
		//return []float64{2.0, 3.0}
		for _, val := range s {
			//fmt.Println(val)
			cur, _ = strconv.ParseFloat(val, 64)
			metrics = append(metrics, cur)
		}
		//}
		fmt.Println(line)
		return metrics
	} else {
		reader.Reset(os.Stdin)
	}
	return []float64{2.0, 3.0} // TODO remove
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
	ch <- prometheus.MustNewConstMetric(collector.ylowMetric, prometheus.CounterValue, getYlow())
	ch <- prometheus.MustNewConstMetric(collector.yavgMetric, prometheus.CounterValue, getYavg())
	ch <- prometheus.MustNewConstMetric(collector.yhighMetric, prometheus.CounterValue, getYhigh())
	ch <- prometheus.MustNewConstMetric(collector.barMetric, prometheus.CounterValue, getYlow())

}

// 22-ymin

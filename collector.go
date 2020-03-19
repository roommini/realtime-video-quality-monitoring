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
type VqmCollector struct {
	yminMetric            *prometheus.Desc
	ylowMetric            *prometheus.Desc
	yavgMetric            *prometheus.Desc
	yhighMetric           *prometheus.Desc
	ymaxMetric            *prometheus.Desc
	uminMetric            *prometheus.Desc
	ulowMetric            *prometheus.Desc
	uavgMetric            *prometheus.Desc
	uhighMetric           *prometheus.Desc
	umaxMetric            *prometheus.Desc
	vminMetric            *prometheus.Desc
	vlowMetric            *prometheus.Desc
	vavgMetric            *prometheus.Desc
	vhighMetric           *prometheus.Desc
	vmaxMetric            *prometheus.Desc
	satminMetric          *prometheus.Desc
	satlowMetric          *prometheus.Desc
	satavgMetric          *prometheus.Desc
	sathighMetric         *prometheus.Desc
	satmaxMetric          *prometheus.Desc
	huemedMetric          *prometheus.Desc
	hueavgMetric          *prometheus.Desc
	ydifMetric            *prometheus.Desc
	udifMetric            *prometheus.Desc
	vdifMetric            *prometheus.Desc
	ybitdepthMetric       *prometheus.Desc
	ubitdepthMetric       *prometheus.Desc
	vbitdepthMetric       *prometheus.Desc
	toutMetric            *prometheus.Desc
	vrepMetric            *prometheus.Desc
	brngMetric            *prometheus.Desc
	luminanceMetric       *prometheus.Desc
	new_luminanceMetric   *prometheus.Desc
	relative_changeMetric *prometheus.Desc
	bitplanenoise01Metric *prometheus.Desc
	bitplanenoise11Metric *prometheus.Desc
	bitplanenoise21Metric *prometheus.Desc
}

//You must create a constructor for you collector that
//initializes every descriptor and returns a pointer to the collector
func newVqmCollector() *VqmCollector {
	return &VqmCollector{
		yminMetric:            prometheus.NewDesc("ymin", "ymin metric", nil, nil),
		ylowMetric:            prometheus.NewDesc("ylow", "ylow metric", nil, nil),
		yavgMetric:            prometheus.NewDesc("yavg", "yavg metric", nil, nil),
		yhighMetric:           prometheus.NewDesc("yhigh", "yhigh metric", nil, nil),
		ymaxMetric:            prometheus.NewDesc("ymax", "ymax metric", nil, nil),
		uminMetric:            prometheus.NewDesc("umin", "umin metric", nil, nil),
		ulowMetric:            prometheus.NewDesc("ulow", "ulow metric", nil, nil),
		uavgMetric:            prometheus.NewDesc("uavg", "uavg metric", nil, nil),
		uhighMetric:           prometheus.NewDesc("uhigh", "uhigh metric", nil, nil),
		umaxMetric:            prometheus.NewDesc("umax", "umax metric", nil, nil),
		vminMetric:            prometheus.NewDesc("vmin", "vmin metric", nil, nil),
		vlowMetric:            prometheus.NewDesc("vlow", "vlow metric", nil, nil),
		vavgMetric:            prometheus.NewDesc("vavg", "vavg metric", nil, nil),
		vhighMetric:           prometheus.NewDesc("vhigh", "vhigh metric", nil, nil),
		vmaxMetric:            prometheus.NewDesc("vmax", "vmax metric", nil, nil),
		satminMetric:          prometheus.NewDesc("satmin", "satmin metric", nil, nil),
		satlowMetric:          prometheus.NewDesc("satlow", "satlow metric", nil, nil),
		satavgMetric:          prometheus.NewDesc("satavg", "satavg metric", nil, nil),
		sathighMetric:         prometheus.NewDesc("sathigh", "sathigh metric", nil, nil),
		satmaxMetric:          prometheus.NewDesc("satmax", "satmax metric", nil, nil),
		huemedMetric:          prometheus.NewDesc("huemed", "huemed metric", nil, nil),
		hueavgMetric:          prometheus.NewDesc("hueavg", "hueavg metric", nil, nil),
		ydifMetric:            prometheus.NewDesc("ydif", "ydif metric", nil, nil),
		udifMetric:            prometheus.NewDesc("udif", "udif metric", nil, nil),
		vdifMetric:            prometheus.NewDesc("vdif", "vdif metric", nil, nil),
		ybitdepthMetric:       prometheus.NewDesc("ybitdepth", "ybitdepth metric", nil, nil),
		ubitdepthMetric:       prometheus.NewDesc("ubitdepth", "ubitdepth metric", nil, nil),
		vbitdepthMetric:       prometheus.NewDesc("vbitdepth", "vbitdepth metric", nil, nil),
		toutMetric:            prometheus.NewDesc("tout", "tout metric", nil, nil),
		vrepMetric:            prometheus.NewDesc("vrep", "vrep metric", nil, nil),
		brngMetric:            prometheus.NewDesc("brng", "brng metric", nil, nil),
		luminanceMetric:       prometheus.NewDesc("luminance", "luminance metric", nil, nil),
		new_luminanceMetric:   prometheus.NewDesc("new_luminance", "new_luminance metric", nil, nil),
		relative_changeMetric: prometheus.NewDesc("relative_change", "relative_change metric", nil, nil),
		bitplanenoise01Metric: prometheus.NewDesc("bitplanenoise01", "bitplanenoise01 metric", nil, nil),
		bitplanenoise11Metric: prometheus.NewDesc("bitplanenoise11", "bitplanenoise11 metric", nil, nil),
		bitplanenoise21Metric: prometheus.NewDesc("bitplanenoise21", "bitplanenoise21 metric", nil, nil),
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
	lst := strings.Split(Line, ",")[29:]
	for _, metric := range lst {
		if n, err := strconv.ParseFloat(metric, 64); err == nil {
			numbers = append(numbers, n)
		}
	}
	return numbers
}

//Each and every collector must implement the Describe function.
//It essentially writes all descriptors to the prometheus desc channel.
func (collector *VqmCollector) Describe(ch chan<- *prometheus.Desc) {

	//Update this section with the each metric you create for a given collector
	ch <- collector.yminMetric
	ch <- collector.ylowMetric
	ch <- collector.yavgMetric
	ch <- collector.yhighMetric
	ch <- collector.ymaxMetric
	ch <- collector.uminMetric
	ch <- collector.ulowMetric
	ch <- collector.uavgMetric
	ch <- collector.uhighMetric
	ch <- collector.umaxMetric
	ch <- collector.vminMetric
	ch <- collector.vlowMetric
	ch <- collector.vavgMetric
	ch <- collector.vhighMetric
	ch <- collector.vmaxMetric
	ch <- collector.satminMetric
	ch <- collector.satlowMetric
	ch <- collector.satavgMetric
	ch <- collector.sathighMetric
	ch <- collector.satmaxMetric
	ch <- collector.huemedMetric
	ch <- collector.hueavgMetric
	ch <- collector.ydifMetric
	ch <- collector.udifMetric
	ch <- collector.vdifMetric
	ch <- collector.ybitdepthMetric
	ch <- collector.ubitdepthMetric
	ch <- collector.vbitdepthMetric
	ch <- collector.toutMetric
	ch <- collector.vrepMetric
	ch <- collector.brngMetric
	ch <- collector.luminanceMetric
	ch <- collector.new_luminanceMetric
	ch <- collector.relative_changeMetric
	ch <- collector.bitplanenoise01Metric
	ch <- collector.bitplanenoise11Metric
	ch <- collector.bitplanenoise21Metric

}

//Collect implements required collect function for all promehteus collectors
func (collector *VqmCollector) Collect(ch chan<- prometheus.Metric) {

	//Implement logic here to determine proper metric value to return to prometheus
	//for each descriptor or call other functions that do so.
	//Write latest value for each metric in the prometheus metric channel.
	//Note that you can pass CounterValue, GaugeValue, or UntypedValue types here.
	line := getLine()
	if line != "empty" {
		metrics := convertLine(line)
		ch <- prometheus.MustNewConstMetric(collector.yminMetric, prometheus.GaugeValue, metrics[0])
		ch <- prometheus.MustNewConstMetric(collector.ylowMetric, prometheus.GaugeValue, metrics[1])
		ch <- prometheus.MustNewConstMetric(collector.yavgMetric, prometheus.GaugeValue, metrics[2])
		ch <- prometheus.MustNewConstMetric(collector.yhighMetric, prometheus.GaugeValue, metrics[3])
		ch <- prometheus.MustNewConstMetric(collector.ymaxMetric, prometheus.GaugeValue, metrics[4])
		ch <- prometheus.MustNewConstMetric(collector.uminMetric, prometheus.GaugeValue, metrics[5])
		ch <- prometheus.MustNewConstMetric(collector.ulowMetric, prometheus.GaugeValue, metrics[6])
		ch <- prometheus.MustNewConstMetric(collector.uavgMetric, prometheus.GaugeValue, metrics[7])
		ch <- prometheus.MustNewConstMetric(collector.uhighMetric, prometheus.GaugeValue, metrics[8])
		ch <- prometheus.MustNewConstMetric(collector.umaxMetric, prometheus.GaugeValue, metrics[9])
		ch <- prometheus.MustNewConstMetric(collector.vminMetric, prometheus.GaugeValue, metrics[10])
		ch <- prometheus.MustNewConstMetric(collector.vlowMetric, prometheus.GaugeValue, metrics[11])
		ch <- prometheus.MustNewConstMetric(collector.vavgMetric, prometheus.GaugeValue, metrics[12])
		ch <- prometheus.MustNewConstMetric(collector.vhighMetric, prometheus.GaugeValue, metrics[13])
		ch <- prometheus.MustNewConstMetric(collector.vmaxMetric, prometheus.GaugeValue, metrics[14])
		ch <- prometheus.MustNewConstMetric(collector.satminMetric, prometheus.GaugeValue, metrics[15])
		ch <- prometheus.MustNewConstMetric(collector.satlowMetric, prometheus.GaugeValue, metrics[16])
		ch <- prometheus.MustNewConstMetric(collector.satavgMetric, prometheus.GaugeValue, metrics[17])
		ch <- prometheus.MustNewConstMetric(collector.sathighMetric, prometheus.GaugeValue, metrics[18])
		ch <- prometheus.MustNewConstMetric(collector.satmaxMetric, prometheus.GaugeValue, metrics[19])
		ch <- prometheus.MustNewConstMetric(collector.huemedMetric, prometheus.GaugeValue, metrics[20])
		ch <- prometheus.MustNewConstMetric(collector.hueavgMetric, prometheus.GaugeValue, metrics[21])
		ch <- prometheus.MustNewConstMetric(collector.ydifMetric, prometheus.GaugeValue, metrics[22])
		ch <- prometheus.MustNewConstMetric(collector.udifMetric, prometheus.GaugeValue, metrics[23])
		ch <- prometheus.MustNewConstMetric(collector.vdifMetric, prometheus.GaugeValue, metrics[24])
		ch <- prometheus.MustNewConstMetric(collector.ybitdepthMetric, prometheus.GaugeValue, metrics[25])
		ch <- prometheus.MustNewConstMetric(collector.ubitdepthMetric, prometheus.GaugeValue, metrics[26])
		ch <- prometheus.MustNewConstMetric(collector.vbitdepthMetric, prometheus.GaugeValue, metrics[27])
		ch <- prometheus.MustNewConstMetric(collector.toutMetric, prometheus.GaugeValue, metrics[28])
		ch <- prometheus.MustNewConstMetric(collector.vrepMetric, prometheus.GaugeValue, metrics[29])
		ch <- prometheus.MustNewConstMetric(collector.brngMetric, prometheus.GaugeValue, metrics[30])
		ch <- prometheus.MustNewConstMetric(collector.luminanceMetric, prometheus.GaugeValue, metrics[31])
		ch <- prometheus.MustNewConstMetric(collector.new_luminanceMetric, prometheus.GaugeValue, metrics[32])
		ch <- prometheus.MustNewConstMetric(collector.relative_changeMetric, prometheus.GaugeValue, metrics[33])
		ch <- prometheus.MustNewConstMetric(collector.bitplanenoise01Metric, prometheus.GaugeValue, metrics[34])
		ch <- prometheus.MustNewConstMetric(collector.bitplanenoise11Metric, prometheus.GaugeValue, metrics[35])
		ch <- prometheus.MustNewConstMetric(collector.bitplanenoise21Metric, prometheus.GaugeValue, metrics[36])

	} else {
		log.Printf("line: %v", line)
	}

}

// 22-ymin

package main

import (
	"fmt"
	"github.com/marcw/rtmetrics"
	"github.com/rcrowley/go-librato"
	"os"
	"time"
)

type libratoCollector struct {
	m librato.Metrics
}

func (lc *libratoCollector) Measure(metrics string, value uint64) {
	lc.m.GetGauge(metrics) <- int64(value)
}

func (lc *libratoCollector) Flush() {
	// NOOP
}

func main() {
	m := librato.NewMetrics(os.Getenv("LIBRATO_USER"), os.Getenv("LIBRATO_TOKEN"), "rtmetrics")
	collector := &libratoCollector{m}
	go rtmetrics.Run(collector, "my-go-instance")

	for _ = range time.Tick(1 * time.Second) {
		fmt.Println("foobar")
	}
}

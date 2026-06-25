package version

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var silaInfo = promauto.NewGauge(prometheus.GaugeOpts{
	Name: "sila_version",
	ConstLabels: prometheus.Labels{
		"version":   gitTag,
		"commit":    gitCommit,
		"buildDate": buildDateUnix},
})

func init() {
	silaInfo.Set(float64(1))
}

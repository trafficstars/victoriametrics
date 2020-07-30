package metrics

import (
	"github.com/VictoriaMetrics/metrics"
)

type MetricGauge struct {
	common
	*metrics.Gauge
}

func GaugeInt64Func(key string, tags AnyTags, fn func() int64) *MetricGauge {
	return &MetricGauge{
		Gauge: metrics.GetOrCreateGauge(prepareKey(key)+"{"+toStringTags(tags)+"}", func() float64 {
			return float64(fn())
		}),
	}
}

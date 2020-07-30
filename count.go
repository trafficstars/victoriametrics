package metrics

import (
	"github.com/VictoriaMetrics/metrics"
)

type MetricCount struct {
	common
	*metrics.Counter
}

func (c *MetricCount) Increment() {
	c.Inc()
}

func (c *MetricCount) Decrement() {
	c.Dec()
}

func Count(key string, tags AnyTags) *MetricCount {
	return &MetricCount{
		Counter: metrics.GetOrCreateCounter(prepareKey(key) + "{" + toStringTags(tags) + "}"),
	}
}

func GaugeInt64(key string, tags AnyTags) *MetricCount {
	return &MetricCount{
		Counter: metrics.GetOrCreateCounter(prepareKey(key) + "{" + toStringTags(tags) + "}"),
	}
}

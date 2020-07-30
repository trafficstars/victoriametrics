package metrics

import (
	"time"

	"github.com/VictoriaMetrics/metrics"
)

type MetricTimingFlow struct {
	common
	*metrics.Summary
}

func (t *MetricTimingFlow) ConsiderValue(d time.Duration) {
	t.Update(float64(d))
}

func TimingFlow(key string, tags AnyTags) *MetricTimingFlow {
	return &MetricTimingFlow{
		Summary: metrics.GetOrCreateSummary(prepareKey(key) + "{" + toStringTags(tags) + "}"),
	}
}

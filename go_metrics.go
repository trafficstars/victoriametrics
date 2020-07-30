package metrics

import (
	"runtime"

	"github.com/VictoriaMetrics/metrics"
)

func CollectGoMetrics() {
	memStats := &runtime.MemStats{}
	runtime.ReadMemStats(memStats)
	metrics.GetOrCreateCounter("num_goroutine{" + defaultTags + "}").Set(uint64(runtime.NumGoroutine()))
	metrics.GetOrCreateCounter("num_cgo_call{" + defaultTags + "}").Set(uint64(runtime.NumCgoCall()))
	metrics.GetOrCreateCounter("num_cpu{" + defaultTags + "}").Set(uint64(runtime.NumCPU()))

	metrics.GetOrCreateCounter("mem_Frees{" + defaultTags + "}").Set(memStats.Frees)
	metrics.GetOrCreateCounter("mem_Mallocs{" + defaultTags + "}").Set(memStats.Mallocs)
	metrics.GetOrCreateCounter("mem_Lookups{" + defaultTags + "}").Set(memStats.Lookups)
	metrics.GetOrCreateCounter("mem_Sys{" + defaultTags + "}").Set(memStats.Sys)
	metrics.GetOrCreateCounter("mem_TotalAlloc{" + defaultTags + "}").Set(memStats.TotalAlloc)
	metrics.GetOrCreateCounter("mem_Alloc{" + defaultTags + "}").Set(memStats.Alloc)
	metrics.GetOrCreateCounter("mem_GCCPUFraction{" + defaultTags + "}").Set(uint64(memStats.GCCPUFraction))
	metrics.GetOrCreateCounter("mem_NumForcedGC{" + defaultTags + "}").Set(uint64(memStats.NumForcedGC))
	metrics.GetOrCreateCounter("mem_NumGC{" + defaultTags + "}").Set(uint64(memStats.NumGC))
	metrics.GetOrCreateCounter("mem_PauseTotalNs{" + defaultTags + "}").Set(memStats.PauseTotalNs)
	metrics.GetOrCreateCounter("mem_LastGC{" + defaultTags + "}").Set(memStats.LastGC)
	metrics.GetOrCreateCounter("mem_NextGC{" + defaultTags + "}").Set(memStats.NextGC)
	metrics.GetOrCreateCounter("mem_OtherSys{" + defaultTags + "}").Set(memStats.OtherSys)
	metrics.GetOrCreateCounter("mem_GCSys{" + defaultTags + "}").Set(memStats.GCSys)
	metrics.GetOrCreateCounter("mem_BuckHashSys{" + defaultTags + "}").Set(memStats.BuckHashSys)
	metrics.GetOrCreateCounter("mem_MCacheSys{" + defaultTags + "}").Set(memStats.MCacheSys)
	metrics.GetOrCreateCounter("mem_MCacheInuse{" + defaultTags + "}").Set(memStats.MCacheInuse)
	metrics.GetOrCreateCounter("mem_MSpanSys{" + defaultTags + "}").Set(memStats.MSpanSys)
	metrics.GetOrCreateCounter("mem_MSpanInuse{" + defaultTags + "}").Set(memStats.MSpanInuse)
	metrics.GetOrCreateCounter("mem_StackSys{" + defaultTags + "}").Set(memStats.StackSys)
	metrics.GetOrCreateCounter("mem_StackInuse{" + defaultTags + "}").Set(memStats.StackInuse)
	metrics.GetOrCreateCounter("mem_HeapObjects{" + defaultTags + "}").Set(memStats.HeapObjects)
	metrics.GetOrCreateCounter("mem_HeapReleased{" + defaultTags + "}").Set(memStats.HeapReleased)
	metrics.GetOrCreateCounter("mem_HeapInuse{" + defaultTags + "}").Set(memStats.HeapInuse)
	metrics.GetOrCreateCounter("mem_HeapIdle{" + defaultTags + "}").Set(memStats.HeapIdle)
	metrics.GetOrCreateCounter("mem_HeapSys{" + defaultTags + "}").Set(memStats.HeapSys)
	metrics.GetOrCreateCounter("mem_HeapAlloc{" + defaultTags + "}").Set(memStats.HeapAlloc)
	metrics.GetOrCreateCounter("mem_BuckHashSys{" + defaultTags + "}").Set(memStats.BuckHashSys)
	metrics.GetOrCreateCounter("mem_BuckHashSys{" + defaultTags + "}").Set(memStats.BuckHashSys)
	metrics.GetOrCreateCounter("mem_BuckHashSys{" + defaultTags + "}").Set(memStats.BuckHashSys)

}

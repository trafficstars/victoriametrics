package metrics

import (
	"bytes"
	"sync"
)

var (
	memoryReuse = true
)

// SetMemoryReuseEnabled defines if memory reuse will be enabled (default -- enabled).
func SetMemoryReuseEnabled(isEnabled bool) {
	memoryReuse = isEnabled
}

type bytesBuffer struct {
	bytes.Buffer
}

type stringSlice []string

func (p stringSlice) Len() int           { return len(p) }
func (p stringSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p stringSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

var (
	bytesBufferPool = &sync.Pool{
		New: func() interface{} {
			return &bytesBuffer{}
		},
	}
	stringSlicePool = &sync.Pool{
		New: func() interface{} {
			return &stringSlice{}
		},
	}
	metricCountPool = &sync.Pool{
		New: func() interface{} {
			return &MetricCount{}
		},
	}
)

func newBytesBuffer() *bytesBuffer {
	return bytesBufferPool.Get().(*bytesBuffer)
}

func (buf *bytesBuffer) Release() {
	if !memoryReuse {
		return
	}
	buf.Reset()
	bytesBufferPool.Put(buf)
}

func newStringSlice() *stringSlice {
	return stringSlicePool.Get().(*stringSlice)
}

func (s *stringSlice) Release() {
	if !memoryReuse {
		return
	}
	*s = (*s)[:0]
	stringSlicePool.Put(s)
}

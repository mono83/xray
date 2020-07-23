package mon

import (
	"runtime"
	"time"

	"github.com/mono83/xray"
)

// StartHealthMonitor starts health monitoring thread, that will send runtime metrics
// every second.
func StartHealthMonitor(r xray.Ray) {
	startedAt := time.Now()

	go func() {
		for {
			time.Sleep(time.Second)

			mem := runtime.MemStats{}
			gor := runtime.NumGoroutine()
			runtime.ReadMemStats(&mem)

			r.Gauge("gcs", int64(mem.NumGC))
			r.Gauge("goroutines", int64(gor))
			r.Gauge("sys.malloc", int64(mem.Mallocs))
			r.Gauge("sys.free", int64(mem.Frees))
			r.Gauge("heap.alloc", int64(mem.HeapAlloc))
			r.Gauge("heap.inuse", int64(mem.HeapInuse))
			r.Gauge("heap.sys", int64(mem.HeapSys))
			r.Gauge("heap.objects", int64(mem.HeapObjects))
			r.Gauge("heap.nextgc", int64(mem.NextGC))
			r.Gauge("uptime", int64(time.Now().Sub(startedAt).Seconds()))
		}
	}()
}

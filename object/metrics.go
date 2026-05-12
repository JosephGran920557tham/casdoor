package object

import (
	"runtime"
	"time"
)

var startTime = time.Now()

// MetricsInfo holds runtime and application metrics
type MetricsInfo struct {
	Uptime        string `json:"uptime"`
	GoRoutines    int    `json:"goRoutines"`
	MemAllocMB    float64 `json:"memAllocMb"`
	MemSysMB      float64 `json:"memSysMb"`
	NumGC         uint32 `json:"numGc"`
	CPUs          int    `json:"cpus"`
	GoVersion     string `json:"goVersion"`
}

// GetMetricsInfo collects and returns current runtime metrics
func GetMetricsInfo() *MetricsInfo {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	uptime := time.Since(startTime)

	return &MetricsInfo{
		Uptime:     formatDuration(uptime),
		GoRoutines: runtime.NumGoroutine(),
		MemAllocMB: bytesToMB(memStats.Alloc),
		MemSysMB:   bytesToMB(memStats.Sys),
		NumGC:      memStats.NumGC,
		CPUs:       runtime.NumCPU(),
		GoVersion:  runtime.Version(),
	}
}

func bytesToMB(b uint64) float64 {
	return float64(b) / 1024 / 1024
}

func formatDuration(d time.Duration) string {
	days := int(d.Hours()) / 24
	hours := int(d.Hours()) % 24
	minutes := int(d.Minutes()) % 60
	seconds := int(d.Seconds()) % 60
	if days > 0 {
		return fmt.Sprintf("%dd%dh%dm%ds", days, hours, minutes, seconds)
	}
	return fmt.Sprintf("%dh%dm%ds", hours, minutes, seconds)
}

package object

import (
	"runtime"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetMetricsInfo_Structure(t *testing.T) {
	metrics := GetMetricsInfo()

	assert.NotNil(t, metrics)
	assert.NotEmpty(t, metrics.Uptime)
	assert.NotEmpty(t, metrics.GoVersion)
	assert.Greater(t, metrics.GoRoutines, 0)
	assert.Greater(t, metrics.CPUs, 0)
	assert.GreaterOrEqual(t, metrics.MemAllocMB, float64(0))
	assert.GreaterOrEqual(t, metrics.MemSysMB, float64(0))
}

func TestGetMetricsInfo_GoVersion(t *testing.T) {
	metrics := GetMetricsInfo()
	assert.Equal(t, runtime.Version(), metrics.GoVersion)
}

func TestGetMetricsInfo_CPUs(t *testing.T) {
	metrics := GetMetricsInfo()
	assert.Equal(t, runtime.NumCPU(), metrics.CPUs)
}

func TestFormatDuration(t *testing.T) {
	tests := []struct {
		duration time.Duration
		expected string
	}{
		{time.Hour*25 + time.Minute*30 + time.Second*15, "1d1h30m15s"},
		{time.Hour*2 + time.Minute*5 + time.Second*3, "2h5m3s"},
		{time.Minute*45 + time.Second*10, "0h45m10s"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			result := formatDuration(tt.duration)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestBytesToMB(t *testing.T) {
	result := bytesToMB(1024 * 1024)
	assert.Equal(t, float64(1), result)

	result = bytesToMB(0)
	assert.Equal(t, float64(0), result)
}

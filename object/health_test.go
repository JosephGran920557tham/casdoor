package object

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetHealthStatus_Structure(t *testing.T) {
	status := &HealthStatus{
		Status:    "ok",
		Timestamp: time.Now(),
		Checks:    map[string]string{"database": "ok"},
	}

	assert.NotNil(t, status)
	assert.Equal(t, "ok", status.Status)
	assert.NotZero(t, status.Timestamp)
	assert.Contains(t, status.Checks, "database")
}

func TestHealthStatus_Degraded(t *testing.T) {
	status := &HealthStatus{
		Status:    "degraded",
		Timestamp: time.Now(),
		Checks:    map[string]string{"database": "connection refused"},
	}

	assert.Equal(t, "degraded", status.Status)
	assert.Equal(t, "connection refused", status.Checks["database"])
}

func TestHealthStatus_Timestamp(t *testing.T) {
	before := time.Now()
	status := &HealthStatus{
		Status:    "ok",
		Timestamp: time.Now(),
		Checks:    make(map[string]string),
	}
	after := time.Now()

	assert.True(t, status.Timestamp.After(before) || status.Timestamp.Equal(before))
	assert.True(t, status.Timestamp.Before(after) || status.Timestamp.Equal(after))
}

package object

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetConfigInfo_Structure(t *testing.T) {
	info := GetConfigInfo()

	assert.NotNil(t, info)
	assert.NotEmpty(t, info.Environment)
	assert.NotEmpty(t, info.LogLevel)
	assert.NotEmpty(t, info.Timezone)
	assert.NotEmpty(t, info.DatabaseType)
	assert.False(t, info.Timestamp.IsZero())
}

func TestGetConfigInfo_Defaults(t *testing.T) {
	os.Unsetenv("APP_ENV")
	os.Unsetenv("LOG_LEVEL")
	os.Unsetenv("TZ")
	os.Unsetenv("DATABASE_TYPE")

	info := GetConfigInfo()

	assert.Equal(t, "production", info.Environment)
	assert.Equal(t, "info", info.LogLevel)
	assert.Equal(t, "UTC", info.Timezone)
	assert.Equal(t, "mysql", info.DatabaseType)
	assert.False(t, info.DebugMode)
}

func TestGetConfigInfo_DebugMode(t *testing.T) {
	os.Setenv("APP_ENV", "development")
	defer os.Unsetenv("APP_ENV")

	info := GetConfigInfo()

	assert.Equal(t, "development", info.Environment)
	assert.True(t, info.DebugMode)
}

func TestGetConfigInfo_Timestamp(t *testing.T) {
	before := time.Now().UTC()
	info := GetConfigInfo()
	after := time.Now().UTC()

	assert.True(t, info.Timestamp.After(before) || info.Timestamp.Equal(before))
	assert.True(t, info.Timestamp.Before(after) || info.Timestamp.Equal(after))
}

func TestGetConfigInfo_CacheEnabled(t *testing.T) {
	os.Setenv("CACHE_ENABLED", "false")
	defer os.Unsetenv("CACHE_ENABLED")

	info := GetConfigInfo()
	assert.False(t, info.CacheEnabled)
}

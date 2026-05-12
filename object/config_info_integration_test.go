package object

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetConfigInfo_CustomEnv(t *testing.T) {
	os.Setenv("APP_ENV", "staging")
	os.Setenv("LOG_LEVEL", "debug")
	os.Setenv("TZ", "America/New_York")
	os.Setenv("DATABASE_TYPE", "postgres")
	defer func() {
		os.Unsetenv("APP_ENV")
		os.Unsetenv("LOG_LEVEL")
		os.Unsetenv("TZ")
		os.Unsetenv("DATABASE_TYPE")
	}()

	info := GetConfigInfo()

	assert.Equal(t, "staging", info.Environment)
	assert.Equal(t, "debug", info.LogLevel)
	assert.Equal(t, "America/New_York", info.Timezone)
	assert.Equal(t, "postgres", info.DatabaseType)
	assert.False(t, info.DebugMode)
}

func TestGetConfigInfo_DevEnvironment(t *testing.T) {
	os.Setenv("APP_ENV", "dev")
	defer os.Unsetenv("APP_ENV")

	info := GetConfigInfo()

	assert.Equal(t, "dev", info.Environment)
	assert.True(t, info.DebugMode)
}

func TestGetConfigInfo_MultipleCallsConsistent(t *testing.T) {
	os.Setenv("APP_ENV", "production")
	defer os.Unsetenv("APP_ENV")

	info1 := GetConfigInfo()
	info2 := GetConfigInfo()

	assert.Equal(t, info1.Environment, info2.Environment)
	assert.Equal(t, info1.DebugMode, info2.DebugMode)
	assert.Equal(t, info1.DatabaseType, info2.DatabaseType)
	assert.Equal(t, info1.LogLevel, info2.LogLevel)
}

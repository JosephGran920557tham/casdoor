package object

import (
	"os"
	"time"
)

// ConfigInfo holds runtime configuration information
type ConfigInfo struct {
	Environment  string    `json:"environment"`
	DebugMode    bool      `json:"debugMode"`
	DatabaseType string    `json:"databaseType"`
	CacheEnabled bool      `json:"cacheEnabled"`
	LogLevel     string    `json:"logLevel"`
	Timezone     string    `json:"timezone"`
	Timestamp    time.Time `json:"timestamp"`
}

// GetConfigInfo returns the current runtime configuration information
func GetConfigInfo() *ConfigInfo {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "production"
	}

	logLevel := os.Getenv("LOG_LEVEL")
	if logLevel == "" {
		logLevel = "info"
	}

	tz := os.Getenv("TZ")
	if tz == "" {
		tz = "UTC"
	}

	dbType := os.Getenv("DATABASE_TYPE")
	if dbType == "" {
		dbType = "mysql"
	}

	debugMode := env == "development" || env == "dev"
	cacheEnabled := os.Getenv("CACHE_ENABLED") != "false"

	return &ConfigInfo{
		Environment:  env,
		DebugMode:    debugMode,
		DatabaseType: dbType,
		CacheEnabled: cacheEnabled,
		LogLevel:     logLevel,
		Timezone:     tz,
		Timestamp:    time.Now().UTC(),
	}
}

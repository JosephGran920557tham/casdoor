package object

import (
	"runtime"
	"time"
)

// BuildInfo holds build-time metadata injected via ldflags
var (
	BuildVersion = "dev"
	BuildCommit  = "unknown"
	BuildDate    = "unknown"
)

// VersionInfo represents the application version and runtime information
type VersionInfo struct {
	Version   string `json:"version"`
	Commit    string `json:"commit"`
	BuildDate string `json:"buildDate"`
	GoVersion string `json:"goVersion"`
	OS        string `json:"os"`
	Arch      string `json:"arch"`
	Uptime    string `json:"uptime"`
}

var startTime = time.Now()

// GetVersionInfo returns the current application version and runtime metadata
func GetVersionInfo() *VersionInfo {
	return &VersionInfo{
		Version:   BuildVersion,
		Commit:    BuildCommit,
		BuildDate: BuildDate,
		GoVersion: runtime.Version(),
		OS:        runtime.GOOS,
		Arch:      runtime.GOARCH,
		Uptime:    time.Since(startTime).Round(time.Second).String(),
	}
}

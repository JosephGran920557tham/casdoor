package object

import (
	"runtime"
	"time"
)

// These variables are set at build time via ldflags.
var (
	BuildVersion = "dev"
	BuildCommit  = "unknown"
	BuildDate    = "unknown"
	BuildBy      = "unknown"
)

// BuildInfo holds information about the current build.
type BuildInfo struct {
	Version    string    `json:"version"`
	Commit     string    `json:"commit"`
	BuildDate  string    `json:"buildDate"`
	BuiltBy    string    `json:"builtBy"`
	GoVersion  string    `json:"goVersion"`
	OS         string    `json:"os"`
	Arch       string    `json:"arch"`
	ReportedAt time.Time `json:"reportedAt"`
}

// GetBuildInfo returns the current build information.
func GetBuildInfo() *BuildInfo {
	return &BuildInfo{
		Version:    BuildVersion,
		Commit:     BuildCommit,
		BuildDate:  BuildDate,
		BuiltBy:    BuildBy,
		GoVersion:  runtime.Version(),
		OS:         runtime.GOOS,
		Arch:       runtime.GOARCH,
		ReportedAt: time.Now().UTC(),
	}
}

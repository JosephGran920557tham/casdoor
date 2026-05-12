package object

import (
	"os"
	"runtime"
	"time"
)

// SystemInfo holds information about the running system environment
type SystemInfo struct {
	Hostname    string `json:"hostname"`
	OS          string `json:"os"`
	Arch        string `json:"arch"`
	CPUs        int    `json:"cpus"`
	GoVersion   string `json:"goVersion"`
	PID         int    `json:"pid"`
	WorkingDir  string `json:"workingDir"`
	Timestamp   string `json:"timestamp"`
}

// GetSystemInfo returns information about the current system environment
func GetSystemInfo() (*SystemInfo, error) {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}

	workingDir, err := os.Getwd()
	if err != nil {
		workingDir = "unknown"
	}

	return &SystemInfo{
		Hostname:   hostname,
		OS:         runtime.GOOS,
		Arch:       runtime.GOARCH,
		CPUs:       runtime.NumCPU(),
		GoVersion:  runtime.Version(),
		PID:        os.Getpid(),
		WorkingDir: workingDir,
		Timestamp:  time.Now().UTC().Format(time.RFC3339),
	}, nil
}

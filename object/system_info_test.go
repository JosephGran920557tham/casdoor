package object

import (
	"os"
	"runtime"
	"strings"
	"testing"
	"time"
)

func TestGetSystemInfo_Structure(t *testing.T) {
	info, err := GetSystemInfo()
	if err != nil {
		t.Fatalf("GetSystemInfo() returned unexpected error: %v", err)
	}
	if info == nil {
		t.Fatal("GetSystemInfo() returned nil")
	}
}

func TestGetSystemInfo_OS(t *testing.T) {
	info, _ := GetSystemInfo()
	if info.OS != runtime.GOOS {
		t.Errorf("expected OS %q, got %q", runtime.GOOS, info.OS)
	}
}

func TestGetSystemInfo_Arch(t *testing.T) {
	info, _ := GetSystemInfo()
	if info.Arch != runtime.GOARCH {
		t.Errorf("expected Arch %q, got %q", runtime.GOARCH, info.Arch)
	}
}

func TestGetSystemInfo_CPUs(t *testing.T) {
	info, _ := GetSystemInfo()
	if info.CPUs <= 0 {
		t.Errorf("expected CPUs > 0, got %d", info.CPUs)
	}
}

func TestGetSystemInfo_PID(t *testing.T) {
	info, _ := GetSystemInfo()
	if info.PID != os.Getpid() {
		t.Errorf("expected PID %d, got %d", os.Getpid(), info.PID)
	}
}

func TestGetSystemInfo_GoVersion(t *testing.T) {
	info, _ := GetSystemInfo()
	if !strings.HasPrefix(info.GoVersion, "go") {
		t.Errorf("expected GoVersion to start with 'go', got %q", info.GoVersion)
	}
}

func TestGetSystemInfo_Timestamp(t *testing.T) {
	before := time.Now().UTC()
	info, _ := GetSystemInfo()
	after := time.Now().UTC()

	parsed, err := time.Parse(time.RFC3339, info.Timestamp)
	if err != nil {
		t.Fatalf("Timestamp %q is not valid RFC3339: %v", info.Timestamp, err)
	}
	if parsed.Before(before) || parsed.After(after) {
		t.Errorf("Timestamp %q is outside expected range", info.Timestamp)
	}
}

package object

import (
	"runtime"
	"strings"
	"testing"
	"time"
)

func TestGetBuildInfo_Structure(t *testing.T) {
	info := GetBuildInfo()
	if info == nil {
		t.Fatal("GetBuildInfo() returned nil")
	}

	if info.Version == "" {
		t.Error("Version should not be empty")
	}
	if info.Commit == "" {
		t.Error("Commit should not be empty")
	}
	if info.GoVersion == "" {
		t.Error("GoVersion should not be empty")
	}
	if info.OS == "" {
		t.Error("OS should not be empty")
	}
	if info.Arch == "" {
		t.Error("Arch should not be empty")
	}
}

func TestGetBuildInfo_GoVersion(t *testing.T) {
	info := GetBuildInfo()
	expected := runtime.Version()
	if info.GoVersion != expected {
		t.Errorf("expected GoVersion %q, got %q", expected, info.GoVersion)
	}
}

func TestGetBuildInfo_GoVersionPrefix(t *testing.T) {
	info := GetBuildInfo()
	if !strings.HasPrefix(info.GoVersion, "go") {
		t.Errorf("GoVersion should start with 'go', got %q", info.GoVersion)
	}
}

func TestGetBuildInfo_Defaults(t *testing.T) {
	info := GetBuildInfo()
	if info.Version != "dev" {
		t.Errorf("default Version should be 'dev', got %q", info.Version)
	}
	if info.Commit != "unknown" {
		t.Errorf("default Commit should be 'unknown', got %q", info.Commit)
	}
}

func TestGetBuildInfo_Timestamp(t *testing.T) {
	before := time.Now().UTC()
	info := GetBuildInfo()
	after := time.Now().UTC()

	if info.ReportedAt.Before(before) || info.ReportedAt.After(after) {
		t.Error("ReportedAt timestamp is outside expected range")
	}
}
